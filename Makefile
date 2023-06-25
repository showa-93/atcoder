.PHONY: init test new solve save current

init:
	@./shell/init.sh
	$(MAKE) open

test:
	@./shell/test.sh

new-%:
	@./shell/new.sh $*
	$(MAKE) open

solve:
	@./shell/solve.sh

save:
	@./shell/save.sh

open:
	@./shell/open.sh
