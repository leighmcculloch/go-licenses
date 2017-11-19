.PHONY: generate data setup

generate:
	cd data && \
		embedfiles -out ../data.go -pkg licenses .

data:
	rm -fr data && \
		mkdir data && \
		curl -sSL https://github.com/OpenSourceOrg/licenses/archive/master.tar.gz | \
		tar --transform 's/.*/\L&/' -xz --directory data licenses-master/texts/plain/ --exclude='*.py' --strip-components=3

setup:
	go get 4d63.com/embedfiles
