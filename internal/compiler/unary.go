package compiler

import (
	"fmt"

	"github.com/rhumb-lang/rhi/internal/ast"
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func (c *Compiler) compileUnary(unary *ast.UnaryExpression) error {
	// Compile Expression
	// Some unary ops (like #) don't have an expression (it's the Label).
	// But UnaryExpression definition is Op + Expr.
	// If Expr is LabelLiteral, it's just the label.

	if unary.Op == ast.OpSignal {
		// #label -> OP_POST label (Implicit receiver)
		if label, ok := unary.Expr.(*ast.LabelLiteral); ok {
			idx := c.makeConstant(mapval.NewText(label.Value))
			c.emit(mapval.OP_LOAD_CONST)
			c.Chunk().WriteByte(byte(c.makeConstant(mapval.NewEmpty())), 0) // Push Empty as receiver
			c.emit(mapval.OP_POST)
			c.Chunk().WriteByte(byte(idx), 0)
			c.Chunk().WriteByte(0, 0) // Arg count 0
			return nil
		}

		// #label(args...)
		if call, ok := unary.Expr.(*ast.CallExpression); ok {
			if label, ok := call.Callee.(*ast.LabelLiteral); ok {
				// Push Implicit Receiver (Empty)
				c.emit(mapval.OP_LOAD_CONST)
				c.Chunk().WriteByte(byte(c.makeConstant(mapval.NewEmpty())), 0)

				// Compile Args
				for _, arg := range call.Args {
					if err := c.compileExpression(arg); err != nil {
						return err
					}
				}

				// Emit OP_POST
				idx := c.makeConstant(mapval.NewText(label.Value))
				c.emit(mapval.OP_POST)
				c.Chunk().WriteByte(byte(idx), 0)
				c.Chunk().WriteByte(byte(len(call.Args)), 0)
				return nil
			}
		}

		// #expr -> Evaluate expr? Architecture says `obj#click`.
		// The syntax `hash label` or `hash string`?
		// `RhumbLexer.g4` has `Hash`. `prefixOp` is `Hash`.
		// So `# expr`.
		// If expr evaluates to a text/key, we can post it?
		// But `OP_POST` takes a CONSTANT index for the name.
		// It implies static names.
		// If dynamic name, we need dynamic POST opcode?
		// Or we assume `Expr` is static label for now.
		return fmt.Errorf("dynamic signal names not supported yet")
	}

	if unary.Op == ast.OpReply {
		// ^label -> OP_INJECT label
		// ^reply(...) is usually a CallExpression?
		// Grammar: `expression OpenParen ...` is invocation.
		// `^label` is UnaryExpression.
		// If used as `^ack`, it's just the label.
		// If used as `^ack(...)`, it parses as `(^ack)(...)`?
		// No, `prefixOp` binds tight?
		// Let's check precedence.
		// Invocation is `expression OpenParen`.
		// Prefix is `prefixOp expression`.
		// `^ack(...)` -> `(^ack)(...)`.
		// `^ack` evaluates to...?
		// Wait, `^ack` is an action. It performs INJECT.
		// If it's `^ack`, it injects empty payload?
		if label, ok := unary.Expr.(*ast.LabelLiteral); ok {
			// Push Signal (Receiver) - Where is it?
			// In a Selector Pattern Action `msg .. ^ack`, the signal is NOT implicitly on stack.
			// The user must provide it?
			// Test 1: `#ping .. ^pong('hello')`.
			// The pattern `#ping` matches. The subject (Signal) is on stack?
			// Selector logic: `pattern .. action`.
			// Match: `target .. action`.
			// If target is literal, we consume subject. Stack empty.
			// If target is `#ping`, it's a structural match?
			// Architecture: `Signals... bubble up... Selector matches...`
			// The Selector sees the Signal object.
			// `Match` consumes it?
			// If I write `#ping`, I match against the Signal name?
			// The Subject is the Signal.
			// So we need to match Signal Name == "ping".
			// And keep Signal on stack if we want to reply to it?
			// But `^pong` needs the Signal.

			// If `unary` is `^pong`, we need the Signal.
			// Is it implicit?
			// Architecture 7.5: "It checks... Zombie Frame...".
			// The VM `opInject` pops the receiver (Signal).
			// So the Signal MUST be on the stack.

			// In `dummy { #ping .. ^pong }`, `#ping` pattern matches.
			// Does it leave the Signal on the stack?
			// `opSelect` (if we had one) or the logic in `compileSelector`.
			// `compileSelector` loads Subject (Signal).
			// `compilePattern` (`#ping`).
			// If `#ping` matches, we proceed to Action.
			// If we `OP_POP` the subject before Action, `^pong` fails.
			// `compileSelector` does `OP_POP` if not `IsConsume`.
			// Wait, `IsConsume` (..) vs `IsPeek` (::).
			// `..` consumes. So Subject is POPPED.
			// Then `^pong` has no receiver!

			// Unless `#ping` pattern BINDS the signal?
			// Test 3: `[#pong(msg)]`. This binds.
			// Test 1: `#ping`.
			// Maybe `#ping` as a pattern means "Match signal name 'ping' AND leave signal on stack?"
			// Or maybe `^pong` implicitly uses "Current Message"? (Like `self`?)

			// Let's look at `compilePattern`.
			// `compilePattern` handles `BinaryExpression`, `LabelLiteral`, `Literals`.
			// It doesn't handle `UnaryExpression` (e.g. `#ping`).
			// I need to implement `compilePattern` for `UnaryExpression`.

			idx := c.makeConstant(mapval.NewText(label.Value))
			// We need the Signal on stack.
			// Use `OP_LOAD_LOC` for `_`? (Slot 0 in selector).
			c.emit(mapval.OP_LOAD_LOC)
			c.Chunk().WriteByte(0, 0) // Load Subject (Signal)

			c.emit(mapval.OP_INJECT)
			c.Chunk().WriteByte(byte(idx), 0)
			c.Chunk().WriteByte(0, 0) // Arg count 0 (for now)
			return nil
		}

		if call, ok := unary.Expr.(*ast.CallExpression); ok {
			// ^label(args...)
			if label, ok := call.Callee.(*ast.LabelLiteral); ok {
				idx := c.makeConstant(mapval.NewText(label.Value))

				// Push Signal (Receiver) First!
				c.emit(mapval.OP_LOAD_LOC)
				c.Chunk().WriteByte(0, 0)

				// Push Payload Args
				for _, arg := range call.Args {
					if err := c.compileExpression(arg); err != nil {
						return err
					}
				}

				c.emit(mapval.OP_INJECT)
				c.Chunk().WriteByte(byte(idx), 0)
				c.Chunk().WriteByte(byte(len(call.Args)), 0)
				return nil
			}
		}

		// Fallback: Anonymous Reply ^(expr)
		
		// Push Signal (Receiver) from Local 0
		c.emit(mapval.OP_LOAD_LOC)
		c.Chunk().WriteByte(0, 0)

		// Compile Payload
		if err := c.compileExpression(unary.Expr); err != nil {
			return err
		}

		idx := c.makeConstant(mapval.NewText(""))
		c.emit(mapval.OP_INJECT)
		c.Chunk().WriteByte(byte(idx), 0)
		c.Chunk().WriteByte(1, 0) // Arg Count 1 (Payload)
		return nil
	}

	// Other unaries
	if err := c.compileExpression(unary.Expr); err != nil {
		return err
	}

	switch unary.Op {
	case ast.OpToNum:
		c.emit(mapval.OP_COERCE_NUM)
	case ast.OpNegNum:
		c.emit(mapval.OP_NUM_NEG)
	case ast.OpNegBool:
		c.emit(mapval.OP_NOT)
	// ... add others
	default:
		return fmt.Errorf("unsupported unary op: %v", unary.Op)
	}
	return nil
}
