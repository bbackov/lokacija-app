import { View, Text, TextInput, Button,TouchableOpacity,Modal } from 'react-native';
import { useEffect, useState } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import * as Device from 'expo-device';
import { fetchStatus } from '../../src/api/fetchStatus';
import stylesMainPage from './stylesMainPage';
import { fetchDodajUredaj } from '../../src/api/fetchDodajUredaj';



export default function Uredaj(props){

    const [trenutniUredaj,setTrenutniUredaj]=useState({
        id_uredaj:"",
        ime_uređaja:"Nepoznato",
        tip_uređaja:"Nepoznato",
        status_uređaja:"Neaktivno"
    })

    const [aktivacija,setAktivacija]=useState(false)

    useEffect(() => {
        const interval = setInterval(async () => {
          try {
            const token = await AsyncStorage.getItem('token');
            const status = await fetchStatus(parseInt(trenutniUredaj.id_uredaj), token);
            setTrenutniUredaj(prev => ({
              ...prev,
              status_uređaja: status || "Neaktivno",
            }));
          } catch (err) {
            console.log("Greška kod dohvaćanja statusa", err);
          }
        }, 200000); 
      
        return () => clearInterval(interval);
      }, [trenutniUredaj.id_uredaj])

      useEffect(() => {
        const inicijalizirajUredaj = async () => {
          const spremljeniId = await AsyncStorage.getItem('id_uredaj');
          const token = await AsyncStorage.getItem('token');
          const ime = Device.deviceName || 'Nepoznato';
          const tip = Device.osName;
          console.log(`${spremljeniId} inicijalizacija`)
      
          if (spremljeniId && token) {
            try {
              const status = await fetchStatus(parseInt(spremljeniId), token);
              console.log(status)
      
              setTrenutniUredaj({
                id_uredaj: spremljeniId,
                ime_uređaja: ime,
                tip_uređaja: tip,
                status_uređaja: status || 'Neaktivno',
              });
      
              props.setID(spremljeniId);
              setAktivacija(true); 
            } catch (err) {
              console.log('Greška pri dohvaćanju statusa:', err);
            }
          }
        };
      
        inicijalizirajUredaj();
      }, []);

    const dodajUredaj = async () => {
        console.log("započela aktivacija")
        try{
            const ime =  Device.deviceName || 'Nepoznato';
            const tip = Device.osName;
            console.log(ime);
            const token=await AsyncStorage.getItem('token');
            const data=await fetchDodajUredaj(token,ime,tip,"Neaktivno");
            console.log("Odgovor s backenda:", data);
            console.log("Uređaj dodan u bazu, ID:", data.id_uredaj);
            await AsyncStorage.setItem('id_uredaj', data.id_uredaj.toString());
            props.setID(data.id_uredaj.toString());
            console.log("ide gas")


            setTrenutniUredaj(prev=>({
                ...prev,
                    id_uredaj:data.id_uredaj,
                    ime_uređaja: ime || 'Nepoznato',
                    tip_uređaja: tip || 'Nepoznato',
                    status_uređaja: "Offline",
                }));
            setAktivacija(true)
        }catch(err){
            alert(err.message || 'Neuspjeh pri povezivanju.');
            console.log("Greška prilikom aktivacije uređaja:", err);
        }
        };

    return(
        
    <View style={stylesMainPage.uredaji}>
        <Text style={stylesMainPage.tekst}>Ime uređaja:{trenutniUredaj.ime_uređaja}</Text>
        <Text style={stylesMainPage.tekst}>Status:{trenutniUredaj.status_uređaja}</Text>
        <Text style={stylesMainPage.tekst}>Operacijski sustav:{trenutniUredaj.tip_uređaja}</Text>
        {!aktivacija&& <TouchableOpacity style={stylesMainPage.uredajiGumb} onPress={dodajUredaj}>
            <Text style={stylesMainPage.uredajiGumbTekst}>Aktiviraj uređaj</Text>
        </TouchableOpacity>}
    </View>

    );


}