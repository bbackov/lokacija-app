import { View, Text, TextInput, Button,TouchableOpacity,Modal } from 'react-native';
import { useState } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { fetchAzuriraj } from '../../src/api/fetchAzuriraj';
import { useEffect } from 'react';
import stylesmodal from './stylesmodal';
import { Ionicons } from '@expo/vector-icons';
import { fetchObrisiProfil } from '../../src/api/fetchObrisiProfil';
import { useRouter } from 'expo-router';


export default function AzurirajModal(props){
    const router = useRouter();
    
    const [ime,setIme]=useState('');
    const [prezime,setPrezime]=useState('');
    const [email,setEmail]=useState('');
    const [testLozinka1,setTestLozinka1]=useState('');
    const [testLozinka2,setTestLozinka2]=useState('');

    

    useEffect(() => {
        setIme(props.imeOld);
        setPrezime(props.prezimeOld);
        setEmail(props.emailOld);
    }, [props.imeOld, props.prezimeOld, props.emailOld]);

    const azuriraj=async()=>{

        if (testLozinka1 !== testLozinka2) {
            alert("Lozinke se ne podudaraju");
            return;
        }
        
        try{
            const token=await AsyncStorage.getItem("token");
            const data=await fetchAzuriraj(ime,prezime,email,testLozinka1,token);
            await AsyncStorage.setItem('token', data.token);
            alert("Uspješno ažuriranje");
            props.onclose();
        }catch(err){
            alert(err.message || "Neuspješno ažuriranje.");
        }
    }

    const obrisi=async()=>{

      try{
        const token=await AsyncStorage.getItem("token");
        const data=await fetchObrisiProfil(token);
        alert("Uspješno Brisanje");
        await AsyncStorage.removeItem("token");
        router.replace("/korisnik/login");
      }catch(err){
        alert(err.message || 'Neuspjeh pri povezivanju.')
      }

    }




    return(
        
<Modal
  visible={props.visible}
  transparent={true}
  animationType="slide"
>
  <View style={stylesmodal.modalOverlay}>
    <View style={stylesmodal.modalContent}>
      
      
      <TouchableOpacity onPress={props.onClose} style={{ alignSelf: 'flex-end' }}>
        <Ionicons name='close' size={24} color="#333" />
      </TouchableOpacity>

      
      <Text style={stylesmodal.tekst}>Trenutno ime: {props.imeOld}</Text>
      <Text style={stylesmodal.tekst}>Ime:</Text>
      <TextInput
        placeholder='Ime'
        value={ime}
        onChangeText={setIme}
        style={stylesmodal.input}
      />

     
      <Text style={stylesmodal.tekst}>Trenutno prezime: {props.prezimeOld}</Text>
      <Text style={stylesmodal.tekst}>Prezime:</Text>
      <TextInput
        placeholder='Prezime'
        value={prezime}
        onChangeText={setPrezime}
        style={stylesmodal.input}
      />

   
      <Text style={stylesmodal.tekst}>Trenutni E-mail: {props.emailOld}</Text>
      <Text style={stylesmodal.tekst}>E-mail:</Text>
      <TextInput
        placeholder='E-mail'
        value={email}
        onChangeText={setEmail}
        style={stylesmodal.input}
      />

      
      <Text style={stylesmodal.tekst}>Lozinka:</Text>
      <TextInput
        placeholder="*********"
        value={testLozinka1}
        onChangeText={setTestLozinka1}
        secureTextEntry
        style={stylesmodal.input}
      />

      <Text style={stylesmodal.tekst}>Ponovi lozinku:</Text>
      <TextInput
        placeholder="*********"
        value={testLozinka2}
        onChangeText={setTestLozinka2}
        secureTextEntry
        style={stylesmodal.input}
      />

      
      <TouchableOpacity style={stylesmodal.gumb} onPress={azuriraj}>
        <Text style={stylesmodal.gumbTekst}>Ažuriraj podatke</Text>
      </TouchableOpacity>

      <TouchableOpacity style={stylesmodal.gumb} onPress={obrisi}>
        <Text style={stylesmodal.gumbTekst}>Obrisi Profil</Text>
      </TouchableOpacity>
    </View>
  </View>
</Modal>

    );


}