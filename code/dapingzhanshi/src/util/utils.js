function manageUserData(params){
    if(Array.isArray(params)){
        for(let i=0; i<params.length; i++){
            let obj = params[i];
            if(Array.isArray(obj.data)){
                obj.data.map(item=>{
                    let title = item.title.substring(0,2);
                    let value = item.value;
                    return {
                        title,
                        value
                    }
                });
                console.log(obj.data);
            }
        }
    }
}
export default {
    manageUserData
}