setup:
	brew install graphviz

graph: setup
	terraform graph > graph.dot | dot -Tsvg > graph.svg

lint:
	terraform fmt
	terraform validate

create-plan:
	terraform plan -out=plan.out

apply-plan:
	terraform apply plan.out
