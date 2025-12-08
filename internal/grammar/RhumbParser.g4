parser grammar RhumbParser;


options {
    tokenVocab=RhumbLexer;
}

document
    : expressions EOF
    ;

expressions
    : terminator*
      expression
    ( terminator
      expression? )*
    ;

fields
    : terminator*
      field
    ( terminator
      field? )*
    ;

patterns
    : terminator*
      pattern
    ( terminator
      pattern? )*
    ;

terminator
    : Semicolon | NL ;


literal
    : FloatingPoint  # rationalNumber
    | date           # dateNumber
    | Zero           # zeroNumber
    | NumberPart     # wholeNumber
    | Key            # keySymbol
    | text           # textSymbol
    | reference      # referenceLiteral
    | Bang           # labelSymbol
    | Label          # labelSymbol
    | TripleUnderscore # emptyValue
    ;


fieldLiteral
    : Zero
    | NumberPart
    | Key
    | text
    | reference
    | Bang
    | Label
    | TripleUnderscore
    | HookLabel
    ;

expression
    : OpenBracket fields? CloseBracket # map
    | OpenCurly libraryResolver Pipe libraryPath+ Pipe (Version | Dash) CloseCurly # library
    | OpenAnglet Dollar CloseAnglet # childRealm      // <$>
    | OpenAnglet Pipe CloseAnglet   # detachedRealm   // <|>
    | OpenCurly patterns? CloseCurly # selector
    | OpenParen expressions? CloseParen # routine
    | <assoc=left> chainExpression # member
    | expression OpenParen terminator* expressions? CloseParen # invocation
    | expression OpenBracket terminator* (expressions | accessOp)? CloseBracket # access
    | expression OpenCurly terminator* patterns? CloseCurly # effect
    | <assoc=right> expression NL* exponentiationOp NL* expression # power
    | <assoc=right> prefixOp expression # prefix
    | <assoc=left> expression NL* multiplicativeOp NL* expression # multiplicative
    | <assoc=left> expression NL* additiveOp NL* expression # additive
    | <assoc=left> expression NL* comparativeOp NL* expression # comparative
    | <assoc=left> expression NL* identityOp NL* expression # identity
    | <assoc=left> expression NL* conjunctiveOp NL* expression # conjunctive
    | <assoc=left> expression NL* disjunctiveOp NL* expression # disjunctive
    | <assoc=right> expression NL* applicativeOp NL* expression # applicative
    | <assoc=right> expression NL* conditionalOp NL* expression # conditional
    | <assoc=right> (chainExpression | fieldLiteral) NL* assignmentOp NL* expression # assignLabel
    | literal # simpleExpression
    ;

libraryResolver
    : Bang
    | Dash
    | Label
    ;

libraryPath
    : libraryPathComponent+
    ;

libraryPathComponent
    : Label
    | FSlash
    | BSlash
    | Dot
    | Dash
    | NumberPart
    | HookLabel
    ;

chainExpression
    : fieldLiteral
    ( chainOp prefixOp? fieldLiteral
    | OpenBracket terminator* (expression | accessOp)? CloseBracket
    | OpenParen terminator* expressions? CloseParen
    )+
    ;


field
    : Dot fieldLiteral                              # prefixAssignMutField
    | Dot At fieldLiteral                           # prefixAssignMutSubField
    | Colon fieldLiteral                            # prefixAssignImmSubField
    | Colon At fieldLiteral                         # prefixAssignImmSubField
    | Ampersand fieldLiteral                        # prefixSlurpSpread
    | <assoc=right> fieldLiteral NL* ColonColon NL* expression    # assignMutField
    | <assoc=right> At fieldLiteral NL* ColonColon NL* expression # assignMutSubField
    | <assoc=right> fieldLiteral NL* DotDot NL* expression        # assignImmField
    | <assoc=right> At fieldLiteral NL* DotDot NL* expression     # assignImmSubField
    | OpenBracket fields? CloseBracket #simpleMapField
    | expression # simpleField
    ;

pattern
    : <assoc=right> expression NL* DotDot NL* expression     # assignBreakingPattern
    | <assoc=right> expression NL* ColonColon NL* expression # assignFallthroughPattern
    | expression # assignDefaultPattern
    ;

accessOp
    : CloseAnglet    # append
    | OpenAnglet     # unshift
    | Hash           # length
    | QMark          # empty
    | At             # allSubfields
    | Star           # allFields
    | Zero           # elements
    | Dot            # freeze
    | Colon          # copy
    | FSlash         # toDate
    | Dollar         # parameters
    | Caret          # constructor
    | Bang           # base
    | Plus           # toNumber
    | Dash           # negateNumber
    | Equal          # toTruth
    | Tilde          # negateTruth
    | Ampersand      # variadic
    | Backtick       # toKey
    ;

applicativeOp
    : DashGreater    # function
    | BangGreater    # method
    | PlusGreater    # letFunction
    ;

comparativeOp
    : GreaterGreater # greaterThan
    | GreaterEqual   # greaterThanOrEqualTo
    | LesserLesser   # lessThan
    | LesserEqual    # lessThanOrEqualTo
    ;

identityOp
    : EqualEqual     # isEqual
    | EqualBSlash    # isInner
    | EqualAt        # isUnder
    | TildeTilde     # notEqual
    | TildeBSlash    # notInner
    | TildeAt        # notUnder
    ;

conjunctiveOp
    : FSlashBSlash   // and
    ;

disjunctiveOp
    : BSlashFSlash    // or
    ;

conditionalOp
    : PipePipe       # pipe
    | QMarkQMark     # default
    | LesserGreater  # foreach
    | PipeGreater    # while
    | EqualGreater   # then
    | TildeGreater   # else
    | CaretEqual     # destructure
    ;

additiveOp
    : PlusPlus       # addition
    | PlusDash       # deviation
    | DashDash       # subtraction
    | AmpAmp         # concatenate
    ;



multiplicativeOp
    : StarStar       # multiplication
    | FSlashFSlash   # rationalDivision
    | PlusFSlash     # wholeDivision
    | DashFSlash     # remainderDivision
    | BangBang       # bind
    ;

exponentiationOp
    : CaretCaret     # exponent
    | CaretFSlash    # rootExtraction
    | Pipe           # range
    | StarCaret      # scientific
    ;


assignmentOp
    : DotEqual       # immutableLabel
    | ColonEqual     # mutableLabel
    ;

prefixOp
    : QMark          # emptyPrefix
    | Dot            # freezePrefix
    | Colon          # copyPrefix
    | Plus           # toNumberPrefix
    | Dash           # negateNumberPrefix
    | Equal          # toTruthPrefix
    | Tilde          # negateTruthPrefix
    | Ampersand      # variadicPrefix
    | Dollar         # argumentPrefix
    | Hash           # signalOutwardPrefix
    | Caret          # signalInwardPrefix
    ;


chainOp
    : BSlash         # nestedField
    | BSlashBSlash   # deeplyNestedField
    | At             # nestedSubfield
    | AtAt           # deeplyNestedSubfield
    | Hash           # signalField       // realm#event
    | Caret          # replyField        // realm^reply
    | Dollar         # proclamationField // realm$state
    ;

datePart
    : NumberPart
    | Dollar Label
    ;

date
    : datePart FSlash datePart FSlash datePart
    ;

text
    : Bang? OpenInterpString (InnerText | interpolation)* CloseInterpString
    | RawText
    ;

interpolation
    : InterpLabel                             #labelInterp
    | EnterRoutineInterp expressions CloseParen  #routineInterp
    | EnterSelectorInterp expressions CloseCurly #selectorInterp
    ;

reference 
    : OpenAnglet Label CloseAnglet                                         #namedRef
    | OpenAnglet text CloseAnglet                                         #computedRef
    | OpenAnglet OpenParen terminator* expressions CloseParen CloseAnglet     #functionRef
    | OpenAnglet OpenBracket terminator* expressions CloseBracket CloseAnglet #computedRef
    | OpenAnglet OpenCurly terminator* patterns? CloseCurly CloseAnglet       #vassalRef
    ;
