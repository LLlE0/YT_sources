const Input = document.getElementById('inp')
const Button = document.getElementById('btn')
const List = document.getElementById('tasks')

let tasks = []

Button.onclick = function() {
    if (Input.value!=''){
        List.innerHTML = ''
     
        tasks.push({val: Input.value, done:false})
        
        for (let i = 0; i<tasks.length; i++){
            List.insertAdjacentHTML('beforeend', 
                                    getHTML(tasks[i], i))
        }
    }
    Input.value=''
    console.log(tasks)
}

function getHTML(item, ind) {
    return `<li class="list-group-item d-flex 
    justify-content-between align-items-center">
            <span style="word-break:break-all;"
            class=
            ${item.done ? 'text-decoration-line-through' : ''}>
            ${item.val}</span>

            <span>
               <span class="btn btn-outline-${item.done ? 'warning' : 'success'}" 
               data-index=${ind} data-type="OK">&check;</span>

               <span class="btn btn-outline-danger" 
               data-type="ERASE" data-index=${ind}>&cross;
               </span>

               <span class="btn btn-outline-dark" 
               data-type="EDIT" data-index=${ind}>✏️
               </span>
            </span>
        </li>`
}

function render() {
List.innerHTML = ''   
    for (let i = 0; i<tasks.length; i++){
    List.insertAdjacentHTML('beforeend', getHTML(tasks[i], i))
    }
    if (tasks.length==0) {
        List.innerHTML =  '<p>SUBSCRIBE!</p>'
    }
}

List.onclick = function(event){
    if (event.target.dataset.index) {
        const index = Number(event.target.dataset.index)
        const type = event.target.dataset.type
        if (type == 'OK') {
            tasks[index].done = !tasks[index].done
            console.log(tasks[index].done)
        } else if (type=='ERASE') {
            tasks.splice(index, 1)
        } else if (type=='EDIT') {
            var z = prompt('Edit the task:', '')
            if (z) tasks[index].val=z
        }
        render()
    }
}