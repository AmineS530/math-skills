NAME=MathSkills.exec

all: $(NAME) tester
	@echo -n "\n"
	@$$(./math-skills > default_results.txt)
	@$$(./$(NAME) data.txt > my_results.txt)
	@echo "\033[1;38;5;155mResults from go app\033[0m" 
	@cat my_results.txt && echo -n "\n"
	@echo "\033[1;38;5;155mSolution\033[0m" 
	@cat default_results.txt && echo -n "\n"
	@echo "\033[1;38;5;155mRunning diff command to compare both outputs\n(It will display something if there is a mismatch)\n\033[0m" 
	diff default_results.txt my_results.txt

$(NAME):
	@if [ ! -f $(NAME) ]; then \
		echo "\033[1;38;5;155mBuilding our go app...\033[0m"  && \
		go build -o $@ main.go  ; \
	fi

tester:
	@if [ ! -f "math-skills" ]; then \
		echo "\033[1;38;5;155mFetching test script...\033[0m" && \
		wget https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip 2>/dev/null && unzip -qq stat-bin-dockerized.zip; \
		mv stat-bin/bin/math-skills ./ &&  rm -fr stat-bin-dockerized.zip stat-bin ;\
	fi

clean:
	@echo "\033[1;38;5;155mDeleting the added files\033[0m"
	@rm -fr $(NAME) math-skills *.txt

re : clean all