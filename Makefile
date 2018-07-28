.PHONY: doc

doc:
	@echo Making doc
	@pandoc doc/requirements.md --latex-engine xelatex \
		-o doc/requirements.pdf
	@pandoc doc/design.md --latex-engine xelatex -o doc/design.pdf
