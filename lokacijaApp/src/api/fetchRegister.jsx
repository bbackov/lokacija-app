

import { BASE_URL } from "./apiConfig";

export async function fetchRegister(ime,prezime,email,lozinka){
    console.log(BASE_URL)
    try{
        const res=await fetch(`${BASE_URL}/registracija`,{
            method:'POST',
            headers:{ 'Content-Type': 'application/json' },
            body: JSON.stringify({ime,prezime,email,lozinka})
        });
    
        console.log("STATUS:", res.status);
        console.log("HEADERS:", res.headers);
        console.log(res.body)
    const data=await res.json()
    console.log(data);

    if(!res.ok){
        throw new Error(data.poruka || 'Neuspješna registracija.');
    }
    return data;
    }catch(err){
        alert(err.message);
        throw err
    }
}