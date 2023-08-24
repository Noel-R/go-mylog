# go-mylog
Logging package for personal use, makes use of log, written in go.

## How to use

This package implements one type, LoggerContext.  
Define a pointer to this type and use the .Setup(filePath) method to initialise.

Uses customised loggers for my purposes.  
Really just testing out go packages etc.

Console()   -> Logs only to console.  
Info()      -> Logs info to text file & console.  
Warn()      -> Logs warning to text file & console.  
Error()     -> Logs error to text file & console.  
Fatal()     -> Logs error to text file & console then exits with code 1.

