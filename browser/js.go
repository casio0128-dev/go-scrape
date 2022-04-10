package browser

func GetXPath() string {
	return `function getXpath(element) {
	  if(element && element.parentNode) {
		var xpath = getXpath(element.parentNode) + '/' + element.tagName;
		var s = [];
		for(var i = 0; i < element.parentNode.childNodes.length; i++) {
		  var e = element.parentNode.childNodes[i];
		  if(e.tagName == element.tagName) {
			s.push(e);
		  }
		}
		if(1 < s.length) {
		  for(var i = 0; i < s.length; i++) {
			if(s[i] === element) {
			  xpath += '[' + (i+1) + ']';
			  break;
			}
		  }
		}
		return xpath.toLowerCase();
	  } else {
		return '';
	  }
	}`
}
