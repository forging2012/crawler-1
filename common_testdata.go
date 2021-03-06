package crawler

const (
	validLinks = 3

	validHTML = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
    <title>title</title>
    <link rel="stylesheet" type="text/css" href="style.css"/>
    <script type="text/javascript" src="script.js"></script>
  </head>
  <body>
  	<div>
   	<a href="http://example.com">Example</a>
   	<a href="http://example.com/2">Example 2</a>
   	<a href="http://example.com/3">Example 3</a>
   	</div>
  </body>
</html>`

	validHTMLWithURLWithoutHTTP = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
    <title>title</title>
    <link rel="stylesheet" type="text/css" href="style.css"/>
    <script type="text/javascript" src="script.js"></script>
  </head>
  <body>
  	<div>
   	<a href="//www.example.com">Example</a>
   	<a href="//www.example.com/2">Example 2</a>
   	<a href="//www.example.com/3">Example 3</a>
   	</div>
  </body>
</html>`

	validHTMLNoURL = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
    <title>title</title>
    <link rel="stylesheet" type="text/css" href="style.css"/>
    <script type="text/javascript" src="script.js"></script>
  </head>
  <body>
  	<div>
	<p>Text</p> <br/>
   	</div>
  </body>
</html>`

	invalidHTML = `<html>
  <head>
    <title>title</title>
  </head>
  <body>
  	<a/> href="http://example.com">Example</a>
   	<div>
	<p>Text</p> <br/>
   	</div>
</html>`

	withJavascriptURL = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
    <title>title</title>
    <link rel="stylesheet" type="text/css" href="style.css"/>
    <script type="text/javascript" src="script.js"></script>
  </head>
  <body>
  	<div>
   	<a href="javascript:void();">Example</a>
   	<a href="http://example.com/1">Example 1</a>
   	<a href="http://example.com/2">Example 2</a>
   	<a href="http://example.com/3">Example 3</a>
   	</div>
  </body>
</html>`
)
