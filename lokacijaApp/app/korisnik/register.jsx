import { View, Text, TextInput, Button,TouchableOpacity } from 'react-native';
import { useState } from 'react';
import styles from './styles';
import { fetchRegister } from '../../src/api/fetchRegister';
import { useRouter } from 'expo-router';




export default function Register(){

    const router = useRouter();
    const [ime,setIme]=useState('')
    const [prezime,setPrezime]=useState('')
    const [email,setEmail]=useState('')
    const [testLozinka1,setTestLozinka1]=useState('')
    const [testLozinka2,setTestLozinka2]=useState('')


    const register=async ()=>{
        
        if (testLozinka1 !== testLozinka2) {
            alert("Lozinke se ne podudaraju");
            return;
        }else if (!ime || !prezime || !email || !testLozinka1) {
            alert("Sva polja su obavezna.");
            return;
        }
       
        try{
            await fetchRegister(ime,prezime,email,testLozinka1);
            alert("Uspješna regisracija")
            router.replace("/korisnik/login")
            return;
        }catch(err){
                alert(err.message || 'Neuspjeh pri povezivanju.');
        }
        
        
    }

    return(
        <View style={styles.container}>
                <Text style={styles.tekst}>Ime:</Text>
                <TextInput
                    placeholder='Ime:'
                    value={ime}
                    onChangeText={setIme}
                    style={styles.input}
                />

                <Text style={styles.tekst}>Prezime:</Text>
                <TextInput
                    placeholder='Prezime:'
                    value={prezime}
                    onChangeText={setPrezime}
                    style={styles.input}
                />

                <Text style={styles.tekst}>E-mail:</Text>
                <TextInput
                    placeholder='E-mail:'
                    value={email}
                    onChangeText={setEmail}
                    style={styles.input}
                />

                <Text style={styles.tekst}>Lozinka:</Text>
                <TextInput
                    placeholder="*********"
                    value={testLozinka1}
                    onChangeText={setTestLozinka1}
                    secureTextEntry
                    style={styles.input}
                />

                <Text style={styles.tekst}>Ponovi lozinku:</Text>
                <TextInput
                    placeholder="*********"
                    value={testLozinka2}
                    onChangeText={setTestLozinka2}
                    secureTextEntry
                    style={styles.input}
                />                     
            

            <TouchableOpacity style={styles.gumb}  onPress={register}>
                <Text style={styles.gumbTekst} >Registriraj se</Text>
            </TouchableOpacity>


        </View>
    )


}