var findIndex = function(arr,tofind) {
    var i, x;
    for (i in arr) {
      x = arr[i];
      if (x.value === tofind) return parseInt(i);
    }
};
  
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


function checkboxChecked() {
    var checks = document.getElementsByClassName("important-checks");

    console.log(checks)
}