.PHONY: doc

hub:
	@echo making hub
	cd cmd/hub && go install

doc:
	@echo making doc
	java -jar plantuml.jar doc/cache_uml.txt
	cd doc; pandoc requirements.md --latex-engine xelatex \
		-o requirements.pdf
	cd doc; pandoc design.md --latex-engine xelatex -o design.pdf
