import { BASE_URL } from "./apiConfig";


export async function fetchDodajUredaj(token,ime_uredaj,tip_uredaj,status_uredaj) {
    console.log("Šaljem JSON:", JSON.stringify({ ime_uredaj, tip_uredaj, status_uredaj }));
    try{
        const res=await fetch(`${BASE_URL}/dodaj_uređaj`,{
            method:'POST',
            headers:{   'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body:JSON.stringify({ime_uredaj,tip_uredaj,status_uredaj})
        });

        const data =await res.json();

        if(!res.ok){
            throw new Error(data.poruka || 'Neuspješno dodavanje uredaja');
        }
        return data.uredaj
    }catch(err){
        throw(err)
    } 
}