.PHONY: init test new solve save current

init:
	@./shell/init.sh

test:
	@./shell/test.sh

new-%:
	@./shell/new.sh $*

solve:
	@./shell/solve.sh

save:
	@./shell/save.sh

cur:
	@./shell/current.sh
