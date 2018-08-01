# v1.0.0 Requirements

1. Should be a website that shows a configurable website directory
2. Each website in the directory should have a name and a link to the website.
3. Directory should come from a config file
4. Directory should be updatable by modifying the config file while the website
   is running
5. Website should be as easy as possible to install and run

# v1.1.0 Requirements
1. Include a favicon if the website at the URL has one for each website in the
   directory
2. Include a copy-config script that accepts a remote host, user, and path to
   copy the config to
3. Include a deploy script that accepts a remote host, user, and working
   directory to start the website on

# v2.0.0 Requirements

1. Add a favicon.
2. Add a header.
3. Separate styles into a stylesheet.
4. Separate the template into a file.
4. Create a directory called 'static' that holds the favicon and styles.
5. Create a directory called 'tmpl' that holds the template.
4. Add a file server that serves files from the static directory.
5. Cache favicons.
6. Update static directory and template directory for minor-version changes
   automatically.
