# Confluence Page Publisher Usage

a simple utility publishing pages to confluence based on a predefined template

## command line usage

- -filelist defines a list of replacements within the template file.

  the format required is string-to-find;replacement any number of find ; replacements can be provided

- -config the name or path to the config file

### Example

`confluencePagePublisher -config="config-file-name" -fieldlist="%VERSION%;1.0.0;%MSG%;Hello World"`

## Config file

the config file is json based. the fields parentPages and labels are optional all other fields are required

```json
{
"user":"{app-user-name}",
"password":"{app-user-pass}",
"tempplateName":"sample-tempplate.html",
"url":"{server-url}",
"pageTitle":"Page %VERSION%",
"spaceKey":"sample-space",
"labels":[
	{
		"prefix":"global",
		"name":"Test Page"
	}
],
"parentPages":[{"id":"{page-id}"}]

}
```

## Template file

the template can be html or any other format accepted by confluence

### Example

```html
<h1>Test Page %VERSION%</h1>
<p>
%MSG%
</p>
```

# Contributing 

want to Contribute? Feel free to create a pull request

# License

GPL

 

