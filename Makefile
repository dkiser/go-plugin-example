GREETER_PLUGIN_DIRS=$(wildcard ./plugins/greeter/*)
CLUBBER_PLUGIN_DIRS=$(wildcard ./plugins/clubber/*)

all: build-plugins

clean: clean-plugins

build-plugins: $(GREETER_PLUGIN_DIRS) $(CLUBBER_PLUGIN_DIRS)

clean-plugins: 
	rm -f ./plugins/built/*

$(GREETER_PLUGIN_DIRS): 
	$(info Greeter plugins at: $(GREETER_PLUGIN_DIRS))
	$(MAKE) -C $@

$(CLUBBER_PLUGIN_DIRS): 
	$(info Clubber plugins at: $(CLUBBER_PLUGIN_DIRS))
	$(MAKE) -C $@

.PHONY: all $(GREETER_PLUGIN_DIRS) $(CLUBBER_PLUGIN_DIRS)