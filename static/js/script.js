var findIndex = function(arr,tofind) {
    var i, x;
    for (i in arr) {
      x = arr[i];
      if (x.value === tofind) return parseInt(i);
    }
};
  
function changeSelector(selectToChange){
    var phone_selector_element = document.getElementById("phone_selector");
    var tablet_selector_element = document.getElementById("tablet_selector");

    if (selectToChange === "phone") {
        if (!phone_selector_element.classList.contains("selector-selected")){
            // console.log("changing to phone")
            document.getElementById("isphone").checked = true;
            document.getElementById("farm-span-text").innerText = "Farm me some phones "
            
            phone_selector_element.classList.remove("selector-unselected");
            phone_selector_element.classList.add("selector-selected");
            
            tablet_selector_element.classList.remove("selector-selected");
            tablet_selector_element.classList.add("selector-unselected");
        }
    } else if (selectToChange === "tablet") {
        if (!tablet_selector_element.classList.contains("selector-selected")){
            // console.log("changing to tablet")
            document.getElementById("isphone").checked = false;
            document.getElementById("farm-span-text").innerText = "Farm me some tablets "

            tablet_selector_element.classList.remove("selector-unselected");
            tablet_selector_element.classList.add("selector-selected");
        
            phone_selector_element.classList.remove("selector-selected");
            phone_selector_element.classList.add("selector-unselected");        
        }
    }
}

function orderChange() {
    //we need to undisable every select to start over
    undisableallselects()

    //we need to get all selects
    var selects = document.getElementsByClassName("order-selects");
    // console.log(selects)

    for (var i=0; i<selects.length;i++){
        //a value has been set other than  "none"
        if (selects[i].value!==undefined&&selects[i].value!=="none"){
            //we need all the other selects
           var selects2 = document.getElementsByClassName("order-selects")
           var indexofselect = selects2.namedItem(selects[i].name)
           if (indexofselect > -1){
               selects2.splice(indexofselect, 1)
            }
            //looping over other selects we want to disable the value which has been already selected
            for (var j=0; j<selects2.length;j++){
            var indextodisable = findIndex(selects2[j].options,selects[i].value)
            selects2[j].options[indextodisable].disabled = true
            }
        }
    }
}

function undisableallselects(){
    var selects = document.getElementsByClassName("order-selects");
    for (var i=0; i<selects.length;i++){
        // console.log(selects[i].options)
        for (var j=0; j<selects[i].options.length;j++){
            // console.log(selects[i].options[j])
            selects[i].options[j].disabled = false;
        }
    }
}

function undisableallchecks(){
    var checks_i = document.getElementsByClassName("check-important");
    for (var i=0; i<checks_i.length;i++){
        checks_i[i].disabled = false
    }
    var checks_o = document.getElementsByClassName("check-okay");
    for (var i=0; i<checks_o.length;i++){
        checks_o[i].disabled = false
    }    
}

function checkChangeEvent() {
    //we need to undisable every select to start over
    undisableallchecks()

    var checks_i = document.getElementsByClassName("check-important");
    var checks_o = document.getElementsByClassName("check-okay");
    for (var i=0; i<checks_i.length;i++){
        if (checks_i[i].checked) {
            for (var i2=0; i2<checks_o.length;i2++){
                if (checks_i[i].value === checks_o[i2].value) {  
                    checks_o[i2].disabled = true
                }
            }
        }
    }
    for (var j=0; j<checks_o.length;j++){
        if (checks_o[j].checked) {
            for (var j2=0; j2<checks_i.length;j2++){
                if (checks_o[j].value === checks_i[j2].value) {  
                    checks_i[j2].disabled = true
                }
            }
        }
    }    
}

function checkboxChecked() {
    var checks = document.getElementsByClassName("important-checks");

    console.log(checks)
}


