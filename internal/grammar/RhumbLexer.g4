lexer grammar RhumbLexer;

TripleUnderscore : '___' ; // empty

AmpAmp       : '&&' ;
Ampersand    : '&'  ;
At           : '@'  ;
AtAt         : '@@' ;
BSlash       : '\\' ;
BSlashBSlash : '\\\\' ;
BSlashFSlash : '\\/' ;
Backtick     : '`'  ;
Bang         : '!'  ;
BangBang     : '!!' ;
BangGreater  : '!>' ;
Caret        : '^'  ;
CaretCaret   : '^^' ;
CaretEqual   : '^=' ;
CaretFSlash  : '^/' ;
Colon        : ':'  ;
ColonColon   : '::' ;
ColonEqual   : ':=' ;
Comma        : ','  ;
CommaDash    : ',-' ;
Dash         : '-'  ;
DashDash     : '--' ; // subtraction
DashFSlash   : '-/' ; // remainder division
DashGreater  : '->' ; // function
Dollar       : '$'  ; // argument
Dot          : '.'  ; // immutable take
DotDash      : '.-' ; // number part
DotDot       : '..' ; // immutable inner field set
DotEqual     : '.=' ; // immutable local label set
Equal        : '='  ; // logical posit
EqualAt      : '=@' ; // has subfield
EqualBSlash  : '=\\' ; // has field
EqualEqual   : '==' ; // equal
EqualGreater : '=>' ; // and then
FSlash       : '/'  ; // ???
FSlashBSlash : '/\\' ; // logical and
FSlashFSlash : '//' ; // rational division
GreaterEqual   : '>=' ; // greater or equal
GreaterGreater : '>>' ; // greater than 
Hash           : '#'  ; // signal outward
LesserEqual    : '<=' ; // less than or equal
LesserGreater  : '<>' ; // for each 
LesserLesser   : '<<' ; // less than
Pipe         : '|'  ; // ranges
PipeGreater  : '|>' ; // while
PipePipe     : '||' ; // pipe
Plus         : '+'  ; // numerical posit
PlusDash     : '+-' ; // standard deviation
PlusFSlash   : '+/' ; // whole division
PlusPlus     : '++' ; // addition
QMark        : '?'  ; // soft invoke

QMarkQMark   : '??' ; // default
Semicolon    : ';'  ; // separator
Star         : '*'  ; // wildcard
StarCaret    : '*^' ; // scientific notation
StarStar     : '**' ; // multiplication
Tilde        : '~'  ; // logical negate
TildeAt      : '~@' ; // doesn't have subfield
TildeBSlash  : '~\\'; // doesn't have field
TildeTilde   : '~~' ; // not equal
TildeGreater : '~>' ; // or else
PlusGreater  : '+>' ; // IIFE

Version
    : ( NumberPart Dot   NumberPart (DotDash   | Dot NumberPart )
      | NumberPart Comma NumberPart (CommaDash | Comma NumberPart ))
    ( Dash (Letter+|NumberPart)+)?
    ( Plus (Letter+|NumberPart)+)?
    ;
FloatingPoint
    : NumberStart (Comma NumberPart)+ (DotDash | Dot NumberPart )
    | NumberStart (Dot NumberPart)+ (CommaDash | Comma NumberPart )
    | NumberPart  Dot NumberPart
    | NumberPart  Comma NumberPart
    ;
Zero                  : '0' ;
NumberPart            : (Zero | Digit)+ ;
NumberStart           : Digit (Zero | Digit)* ;
fragment Digit        : [1-9] ;

RawText        : '\'' (RawEscape | ~['\n])* '\'' ;
fragment RawEscape : '`\'' | '``' ;
OpenInterpString : '"' -> pushMode(STRING_MODE);

OpenParen    : '('  -> pushMode(DEFAULT_MODE);
CloseParen   : ')'  -> popMode;
OpenBracket  : '['  ;
CloseBracket : ']'  ;
OpenCurly    : '{'  -> pushMode(DEFAULT_MODE);
CloseCurly   : '}'  -> popMode;
OpenAnglet   : '<'  ;
CloseAnglet  : '>'  ;

NL               : '\r'? '\n' | '\r' ;
WS               : [ \t]+ -> skip;
BlockComment     : '%(' ( BlockComment | . )*? '%)' -> channel(HIDDEN) ;
LineComment      : '%' (~[(\r\n] ~[\r\n]*)?   -> channel(HIDDEN) ;

Key
    : Backtick ~[()[\]{} \t\n\r]+
    ;

// Hook Labels for Operator Overloading (Explicit Set)
// Matches _++_, _0_, _>_, _`_, _==_, etc.
// Enclosed in underscores.
// Allowed internal chars: Digits, and ! # $ % & * + , - . / : < = > ? @ ^ | ~ ` \
HookLabel
    : '_' [0-9!#$%&*+,\-./:<=>?@^|~`\\]+ '_'
    ;

Label
    : LabelStart ((LabelPart+ | LabelSymbol LabelPart+)* LabelPart)?
    ;

// Fragment: A Letter is any Unicode Letter, Number-Letter, or Symbol (Emoji).
// We use Unicode Properties (\p) which are mutually exclusive, preventing
// the "used multiple times" error.
fragment Letter
    : [\p{L}]        // Unicode Letters (Latin, Greek, Cyrillic, CJK, etc.)
    | [\p{Mn}]       // Nonspacing Marks (Diacritics & emoji-forcing variation selector-16)
    | [\p{Nl}]       // Number Letters (e.g. Roman Numerals)
    | [\p{So}]       // Symbol Other (Includes most Emojis 🚀, 🧮, 🐚)
    | [\u200D]       // Zero-Width Joiner (Essential for complex Emojis)
    ;
fragment LabelStart  : Star | '_' | Letter ;
fragment LabelPart   : LabelStart | '-' | '.' | Zero | Digit ;
fragment LabelSymbol : Dash | Dot ;


mode STRING_MODE;

EnterRoutineInterp  : '$(' -> pushMode(DEFAULT_MODE);
EnterSelectorInterp : '${' -> pushMode(DEFAULT_MODE);
InterpLabel         : Dollar Label ;

// Covers `n `r `t `b `f `` `$ `"
StringEscape
    : Backtick [nrtbf`$"]
    ;

// InnerText must exclude the Backtick (`)
// so that the StringEscape rule can catch it.
InnerText
    : ~[$\n"`]+
    ;

CloseInterpString   : '"' -> popMode;

