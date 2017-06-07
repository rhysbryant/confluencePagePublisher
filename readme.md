# Confluence Page Publisher 

Is a simple utility publishing pages to confluence based on a predefined template.

# Using it

## Command line Arguments

- -fieldlist defines a list of replacements within the template file.
  The format required is string-to-find;replacement[;...] any number of find ; replacements can be provided.

- -config the name or path to the config file.

- -labels space seperated list of labels

### Example

`confluencePagePublisher -config="config-file-name" -fieldlist="%VERSION%;1.0.0;%MSG%;Hello World" -labels="test v1.0.0"`

## Config file

The config file is json based. the fields parentPages and labels are optional all other fields are required.
Note: the replacements strings for the template (see above) can also be used in some fields in the config file.

**Note:** if the username  and password are missing from the Config file the user will be prompted.

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

The template can be html or any other format accepted by confluence.

### Example

```html
<h1>Test Page %VERSION%</h1>
<p>
%MSG%
</p>
```

# Contributing 

Want to Contribute? Feel free to create a pull request.

# License

GPL

 

