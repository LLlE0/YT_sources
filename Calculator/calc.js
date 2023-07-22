const Input1 = document.getElementById('input1')
const Input2 = document.getElementById('input2')
const PlusBtn = document.getElementById('plus')
const MinusBtn = document.getElementById('minus')
const MultBtn = document.getElementById('mult')
const DivBtn = document.getElementById('div')
const SwapBtn = document.getElementById('swap')
const EqBtn = document.getElementById('eq')
const Result = document.getElementById('output')

let char = '&'

PlusBtn.onclick = function() {
    char='+'
}
MinusBtn.onclick = function() {
    char='-'
}
MultBtn.onclick = function() {
    char='*'
}
DivBtn.onclick = function() {
    char='/'
}
SwapBtn.onclick = function(){
    let Z = Input1.value
    Input1.value=Input2.value
    Input2.value=Z
}
EqBtn.onclick = function(){
    PrintResult()
}

function PrintResult(){
    if (char=='-') {
        Result.value=Number(Input1.value) - Number(Input2.value)
    } else if (char=='+') {
        Result.value=Number(Input1.value) + Number(Input2.value)
    } else if (char=='*') {
        Result.value=Number(Input1.value) * Number(Input2.value)
    } else if (char=='/') {
        Result.value=Number(Input1.value) / Number(Input2.value)
    }
    
}
