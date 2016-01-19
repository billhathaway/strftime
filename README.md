strftime
--
The standard C strftime(3) library is used to format date/timestamps using conversion specifiers, similar to printf/scanf.  

The Go [time](https://golang.org/pkg/time) package uses its own syntax for [time.Format()](https://golang.org/pkg/time/#Time.Format) and [time.Parse()](https://golang.org/pkg/time/#Time.Parse) calls.  This package lets you use a strftime format and get back a string that can be passed to time.Format or time.Parse functions.

Usage Example
--

    // Assuming you want to print out in YYYY-MM-DD HH:MM:SS format 
    format,err := strftime.New("%Y-%m-%d %H:%M:%S")  
    if err != nil {  
      log.Fatal("invalid format")  
    }  
    fmt.Println(time.Now().Format(format))  
