import { View, Text, TextInput, Button ,TouchableOpacity} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import stylesMainPage from './stylesMainPage';
import { useState } from 'react';
import OptionsMenu from './optionsmenu';
import {fetchGetKorisnik} from '../../src/api/fetchGetKorisnik';
import AsyncStorage from '@react-native-async-storage/async-storage';
import AzurirajModal from './azurirajmodal';
import { useEffect } from 'react';
import Uredaj from './uredaj';
import Karta from '../../lokacija/karta';



export default function Home(){

    const [showOption,setOption]=useState(false);
    const [showModal,setModal]=useState(false);
    const [korisnik, setKorisnik] = useState({
        id_korisnik: null,
        ime: '',
        prezime: '',
        email: ''
      });

    const [idUredaj,setUredaj]=useState("")


      useEffect(() => {
        const dohvatiPodatke = async () => {
          try {
            const token = await AsyncStorage.getItem('token');
            const data = await fetchGetKorisnik(token);
            console.log(token);
            setKorisnik(data);
          } catch (err) {
            alert(err.message);
          }
        };
      
        dohvatiPodatke();
      }, []);

    const openClose =()=>{
        setOption(prev=>{
        console.log("Promjena iz", prev, "u", !prev);
        return !prev;
    });
    }

    const Modal=()=>{
        setModal(prev=>!prev);
        console.log("pali modal");
    }

    return(
        <View style={stylesMainPage.container}>
            <View style={stylesMainPage.taskbar}>
                <Text style={stylesMainPage.taskbarTekst}>{korisnik.ime+" "+korisnik.prezime}</Text>
                <View style={{position:"relative" }}> 
                    <Ionicons name="settings-outline" size={30} color="#2f80ed" style={stylesMainPage.ikona} onPress={openClose}/>
                    <OptionsMenu visible={showOption} onClose={()=>setOption(prev=>!prev)} onOpenModal={Modal}/>
                </View>
            </View>

            <Uredaj setID={setUredaj}/>


            <Karta/>
            
            <AzurirajModal
                visible={showModal}
                onClose={Modal}
                imeOld={korisnik.ime}
                prezimeOld={korisnik.prezime}
                emailOld={korisnik.email}
            />

        </View>
    );

}