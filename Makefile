.PHONY: doc

hub:
	@echo making hub
	cd cmd/hub && go install

doc:
	@echo making doc
	pandoc doc/requirements.md --latex-engine xelatex \
		-o doc/requirements.pdf
	pandoc doc/design.md --latex-engine xelatex -o doc/design.pdf
