init:
	@./shell/init.sh

test:
	@./shell/test.sh

new-%:
	@./shell/new.sh $*

save:
	@./shell/save.sh

current:
	@./shell/current.sh
