hello
=====

<pre>
    /get
    /put
    /post
    /delete
    /Sqrt
</pre>

ECMAscript
==========

Acquiring/aspiring/looking at style within this ECMAscript thing

var c = (function () {
    var thePrivate = 3;
    return {
    	get: function () {
	},
	put: function () {
	},
	post: function () {
	},
	delete: function () {
	}
    };
    var thePublic = "private privilleged public"
    var that = this;
    this.please = function (req) { // privilleged
    	...
    }
})();

alert(c.please())
