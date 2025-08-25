

import { BASE_URL } from "./apiConfig";

export async function fetchLogin(email,lozinka){

    try{
        const res=await fetch(`${BASE_URL}/prijava`,{
            method:'POST',
            headers:{ 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, lozinka })
        });
    

    
    const data = await res.json();

    if(!res.ok){
        throw new Error(data.message || 'Neuspješna prijava.');
    }
    console.log("Token iz backenda:", data.token);
    return data;
    
    }catch(err){
        throw(err);
    }
    
}



