import { BASE_URL } from "./apiConfig";



export async function fetchAzuriraj(ime,prezime,email,lozinka,token){

    try{
        const res=await fetch(`${BASE_URL}/azuriraj_profil`,{
            method:'PATCH',
            headers:{   'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    },
            body:JSON.stringify({ime,prezime,email,lozinka})
        });

        const data=await res.json()

        if(!res.ok){
            throw new Error(data.poruka || 'Neuspješno ažuriranje.');
        }
        return data;
    }catch(err){
        throw(err)
    }

}