# v5.1.0 Requirements

1. Use Go modules for versioning

# v5.0.0 Requirements

1. Use cache v2.0.0

# v4.0.0 Requirements

1. Make `cache.DefaultTimeCache` and `cache.DefaultModifiedCache` thread-safe

# v3.0.0 Requirements

1. Fix absolute path handling for favicons not at '/'
2. Update text and background colors to match the favicon
3. Cache the parsed config so that its only reparsed when modified
4. Change the caching for the template so that its only reparsed when modified
5. Log whenever cache items are removed
6. Center containers on the website
7. Don't show anything for favicon if link doesn't exist

# v2.1.0 Requirements

1. Add a copy assets script.

# v2.0.0 Requirements

1. Add a favicon
2. Add a header
3. Separate styles into a stylesheet
4. Separate the template into a file
5. Create a directory called 'static' that holds the favicon and styles
6. Create a directory called 'tmpl' that holds the template
7. Add a file server that serves files from the static directory
8. Cache favicons

# v1.1.0 Requirements
1. Include a favicon if the website at the URL has one for each website in the
   directory
2. Include a copy-config script that accepts a remote host, user, and path to
   copy the config to
3. Include a deploy script that accepts a remote host, user, and working
   directory to start the website on

# v1.0.0 Requirements

1. Should be a website that shows a configurable website directory
2. Each website in the directory should have a name and a link to the website.
3. Directory should come from a config file
4. Directory should be updatable by modifying the config file while the website
   is running
5. Website should be as easy as possible to install and run
