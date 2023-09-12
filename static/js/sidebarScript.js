function ActivateSidebarButton(element) {
    if (element.getAttribute("hx-disable") === "") {
        return;
    }

    element.dispatchEvent(new Event("triggerSwap"));

    var sidebarHeader = document.getElementById("sidebar-placeholder");

    var current = sidebarHeader.getElementsByClassName("active");
    current[0].removeAttribute("hx-disable");
    current[0].className = current[0].className.replace(" active", "");

    element.setAttribute("hx-disable", "");
    element.className += " active";
}

