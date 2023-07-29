const Input1 = document.getElementById('input1')
const Input2 = document.getElementById('input2')
const Current1 = document.getElementById('btn1')
const Current2 = document.getElementById('btn2') 
const Submit = document.getElementById('submit')
const Result = document.getElementById('result')


Current1.onclick = function() {currentTime(Input1)}
Current2.onclick = function() {currentTime(Input2)}

function currentTime(Input){
    const dateTime = new Date()
     
    Input.value= (
    dateTime.getFullYear()+'-'+
    ('0'+dateTime.getMonth()).slice(-2)+'-'+
    dateTime.getDate()
    )

    console.log('click')
}


Submit.onclick = function() {
    let Z = new Date(Input1.value)
    let V = new Date(Input2.value)
    //minutes, hours, days, weeks
    var arr = [
              (V-Z)/(60*1000), 
              (V-Z) / (3600 * 1000), 
              (Math.floor((V-Z)/(24*3600*1000))), 
    Math.round(parseFloat(((V-Z)/(24*3600*1000*7))*100))/100
            ]
    
        
    Render(arr.map((p)=>Math.abs(p)))
}

function Render(arr) {
    Result.innerHTML=`<p>
    <br />Number of <b>minutes</b>: ${arr[0]}
    <br />Number of <b>hours</b>: ${arr[1]}
    <br />Number of <b>days</b>: ${arr[2]}
    <br />Number of <b>weeks</b>: ${arr[3]}
    </p>
    `
}