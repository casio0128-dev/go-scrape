package browser

import (
	"fmt"
	"strings"
)

func MakeScript(scripts ...string) string {
	return strings.Join(scripts, ";")
}

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

func Post() string {
	return `function postXPath(url, act, path, content) {
			const req = {
				action: act,
				target: path,
				content: content,
				currentHref: location.href
			}

			var request = new XMLHttpRequest();
			request.open('POST', url);
			request.onreadystatechange = function () {
				if (request.readyState != 4) {
					// リクエスト中
					console.log("requesting now.");
				} else if (request.status != 200) {
					// 失敗
					console.log("requesting failure.");
				} else {
				// 送信成功
					var result = request.responseText;
					console.log("requesting success.");
					console.log(result);
				}
			};
			request.setRequestHeader('Content-Type', 'application/json');
			console.log(req);
			request.send(JSON.stringify(req));
		}
	`
}

func SetEventListener(url string) string {
	return fmt.Sprintf(`
		document.addEventListener('click', function(e){
			postXPath("%s", 'click', getXpath(e.target), '',);
		});
		document.addEventListener('input', function(e){
			postXPath("%[1]s", 'input', getXpath(e.target), e.target.value);
		});

		let selects = document.getElementsByTagName("select");
		for (let i = 0; i < selects.length; i++) {
			selects[i].addEventListener('change', function(e){
				postXPath("%[1]s", 'select', getXpath(e.target), e.target.value);
			});
		}
	`, url)
}
