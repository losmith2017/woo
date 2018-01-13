ifndef VENDOR_SRCROOT
    $(error VENDOR_SRCROOT not set)
endif

GOPATH             := $(shell pwd)
PORT                = 1380
DESTDIR             = build
VIEW_SRCROOT        = view

VIEW_FILES.src      = index.html
VIEW_FILES.src     += cosmo50.png
VIEW_FILES          = $(VIEW_FILES.src:%=build/%)

VENDOR_FILES.src    = css/bootstrap.min.css
VENDOR_FILES.src   += font-awesome/css/font-awesome.css

# Toastr style 
VENDOR_FILES.src   += css/plugins/toastr/toastr.min.css

# Gritter 
VENDOR_FILES.src   += js/plugins/gritter/jquery.gritter.css

VENDOR_FILES.src   += css/animate.css
VENDOR_FILES.src   += css/style.css

VENDOR_FILES.src   += js/jquery-3.1.1.min.js
VENDOR_FILES.src   += js/bootstrap.min.js
VENDOR_FILES.src   += js/plugins/metisMenu/jquery.metisMenu.js
VENDOR_FILES.src   += js/plugins/slimscroll/jquery.slimscroll.min.js

# Flot
VENDOR_FILES.src   += js/plugins/flot/jquery.flot.js
VENDOR_FILES.src   += js/plugins/flot/jquery.flot.tooltip.min.js
VENDOR_FILES.src   += js/plugins/flot/jquery.flot.spline.js
VENDOR_FILES.src   += js/plugins/flot/jquery.flot.resize.js
VENDOR_FILES.src   += js/plugins/flot/jquery.flot.pie.js

# Peity 
VENDOR_FILES.src   += js/plugins/peity/jquery.peity.min.js
VENDOR_FILES.src   += js/demo/peity-demo.js

# Custom and plugin javascript 
VENDOR_FILES.src   += js/inspinia.js
VENDOR_FILES.src   += js/plugins/pace/pace.min.js

# jQuery UI 
VENDOR_FILES.src   += js/plugins/jquery-ui/jquery-ui.min.js

# GITTER 
VENDOR_FILES.src   += js/plugins/gritter/jquery.gritter.min.js

# Sparkline 
VENDOR_FILES.src   += js/plugins/sparkline/jquery.sparkline.min.js

# Sparkline demo data  
VENDOR_FILES.src   += js/demo/sparkline-demo.js

# ChartJS
VENDOR_FILES.src   += js/plugins/chartJs/Chart.min.js

# Toastr 
VENDOR_FILES.src   += js/plugins/toastr/toastr.min.js

# Images
VENDOR_FILES.src   += css/patterns/header-profile.png
VENDOR_FILES.src   += font-awesome/fonts/fontawesome-webfont.woff
VENDOR_FILES.src   += font-awesome/fonts/fontawesome-webfont.woff2
VENDOR_FILES.src   += font-awesome/fonts/fontawesome-webfont.ttf


VENDOR_FILES        = $(VENDOR_FILES.src:%=build/%)


view-files: $(VIEW_FILES)
vendor-files: $(VENDOR_FILES)

open-browser: 
	open -a "Google Chrome" http://localhost:$(PORT) --args --auto-open-devtools-for-tabs

start: vendor-files
	go run portal.go $(PORT) $(DESTDIR) &
	@sleep 3 && make open-browser

stop:
	kill -HUP $(shell lsof -t -i:$(PORT))

# Order is important. Pick up local first
$(DESTDIR)/%: $(VIEW_SRCROOT)/%
	@mkdir -p $(@D)
	@cp -v $< $(@)

$(DESTDIR)/%: $(VENDOR_SRCROOT)/%
	@mkdir -p $(@D)
	@cp -v $< $(@)

