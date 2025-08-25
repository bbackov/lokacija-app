import { BASE_URL } from "./apiConfig";

export async function fetchStatus(id_uredaja, token) {

    try{
        const res=await fetch(`${BASE_URL}/dohvati_status?id=${id_uredaja}`,{
            method:'GET',
            headers:{   'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
        });

        const data=await res.json()

        if(!res.ok){
            throw new Error(data.poruka || 'Neuspješno dohvaćanje statusa');
        }
        return data.status
    }catch(err){
        throw(err)
    }
    
}