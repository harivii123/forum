function simpleSearch() {
    var input = document.getElementById('author-search');
    var filter = input.value.toUpperCase();
    var ul = document.getElementById("author-list");
    var li = ul.getElementsByTagName("li")

    for (i = 0; i < li.length; i++) {
        var name = li[i].dataset.name;

        if (name.toUpperCase().indexOf(filter) > -1) {
            li[i].style.display = "";
        } else {
            li[i].style.display = "none";
        }
    }
}