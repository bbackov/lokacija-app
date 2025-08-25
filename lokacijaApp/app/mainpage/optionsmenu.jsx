import { View, Text, TextInput, Button ,TouchableOpacity} from 'react-native';
import stylesMainPage from './stylesMainPage';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { fetchOdjava } from '../../src/api/fetchOdjava';
import {useRouter } from 'expo-router';
import { fetchObrisiUredaj } from '../../src/api/fetchObrisiUredaj';

export default function OptionsMenu(props){
    const router = useRouter();
    const tokenErase=async ()=>{
        try {
            const token = await AsyncStorage.getItem("token");
            const id = await AsyncStorage.getItem("id_uredaj");
          
            
            try {
              await fetchOdjava(token);
              await fetchObrisiUredaj(token, id);
            } catch (e) {
              console.log("Odjava nije uspjela na backendus  brišem lokalno:", e.message);
            }
          
            
            await AsyncStorage.removeItem("token");
            await AsyncStorage.removeItem("id_uredaj");
            alert("Odjava zavrsena");
            router.replace("/korisnik/login");
          
          } catch (err) {
            alert(err.message || "Greška pri odjavi");
          }
          
          props.onClose();
    }

    if (props.visible===false){

        return null;
    }else{ 

        return(
            

          
            <View style={stylesMainPage.menuConteiner}>
                <TouchableOpacity style={stylesMainPage.menuItems} onPress={props.onOpenModal}>
                    <Text style={stylesMainPage.menuText} >Ažuriraj profil</Text>
                </TouchableOpacity>
                <TouchableOpacity style={stylesMainPage.menuItems} onPress={tokenErase}>
                    <Text style={stylesMainPage.menuText}>Odjava</Text>
                </TouchableOpacity>
            </View>
        );
    }

}