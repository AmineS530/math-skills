NAME=MathSkills.exec

all: $(NAME) tester
	@$$(./math-skills > default_results.txt)
	@$$(./MathSkills.exec > my_results.txt)
	diff default_results.txt my_results.txt

$(NAME):
	@if [ ! -f $(NAME) ]; then \
	echo "\033[1;38;5;155mBuilding our go app...\033[0m"
	go build -o $@ main.go\
	fi
tester:
	@echo "\033[1;38;5;155mFetching test script...\033[0m"
	@if [ ! -f "math-skills" ]; then \
	wget https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip 2>/dev/null && unzip -qq stat-bin-dockerized.zip; \
	mv stat-bin/bin/math-skills ./ &&  rm -fr stat-bin-dockerized.zip stat-bin ;\
	fi

clean:
	@echo "\033[1;38;5;155mDeleting old builds\033[0m"
	@rm -fr $(NAME) math-skills *.txt