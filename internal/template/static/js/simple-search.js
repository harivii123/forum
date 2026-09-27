function simpleSearch(value1, value2) {
    var filter = value1.value.toUpperCase();
    var ul = document.getElementById(value2);
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