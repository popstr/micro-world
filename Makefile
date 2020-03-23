
gen:
	cd email-service && go generate && cd ..
	cd name-service && go generate && cd ..

