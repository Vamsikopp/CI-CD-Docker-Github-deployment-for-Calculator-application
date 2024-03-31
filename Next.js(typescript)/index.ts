import {question} from 'readline-sync';



const sum = (a:number,b:number) =>{
    let result = a + b;
    return console.log(result);
}
const mines = (a:number,b:number) =>{
    let result = a - b;
    return console.log(result)
}
const divided = (a:number,b:number) =>{
    let result = Math.floor(a / b);
    return console.log(result)
}
const multy = (a:number,b:number) => {
    let result = a * b;
    return console.log(result)
}
const main = () =>{
    const firstNumber :string = question('Enter your first number:\n')
    const secondNumber: string =question('Enter your second number:\n')
    const operator :string = question('Choice your operator one of sum,min,divid,multy:\n')    
    
   if(operator =="sum"){
       sum(+firstNumber,+secondNumber)
   }else{
       console.log('fuck you!')
   }
    

}
main()