
# Themes and _mkpage_

Most content management systems have or evolve a theme engine.
_mkpage_ doesn't have one in the traditional sense but it is
easy to organize your own.  Think about what makes of a theme
engine.

1. a template system
2. layout (hopefully semantic)
3. associated CSS and media assets

Most content management systems impose a certain structure and
approach. Drupal has it's way, Wordpress anther and vernerable
systems like Blogger, Vinette, Moveable Type yet other impositions

Part of the reason these systems impose an approach is because they
need build things to map to their internal structure of their rendering
process.  In a deconstructed content system like _mkpage_ the 
structure is extern.  The build structure is imposed by either a 
Bash script, Makefile or similar mechanism. So implementing themes
in _mkpage_ are flexible and can easily evolve based on the requirements
of the website or service. A theme, for the purposes of _mkpage_,
then is a Bash script describing about to assemble cotent, one or 
more templates to assemble pages from and related CSS and media
documents for rendering something nice in the web browser.

## Demonstration themes examples

+ [simple](simple/)


