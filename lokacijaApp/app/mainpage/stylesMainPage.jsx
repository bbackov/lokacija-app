
import { StyleSheet, Dimensions } from 'react-native';
const { width, height } = Dimensions.get('window');

export default StyleSheet.create({


  container: {
    flex: 1,
    backgroundColor: '#e5e8ea', 
  },

    
    taskbar: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: width *0.02,
        margin:height*0.005,
        borderColor: '#333',           
        borderWidth: width*0.0005,                 
        borderRadius: 12,
        backgroundColor:"#dfe6e9"
        
      },

      taskbarTekst:{
        color:'#333',
        fontSize:22,
        fontWeight:'700',
        margin:height*0.008
      },

      ikona:{
        margin:height*0.008
      },

      uredaji:{
        flexDirection: 'column',      
        alignItems: 'flex-start',       
        justifyContent: 'space-between',
        paddingHorizontal: width *0.03,
        paddingVertical: height*0.05,
        margin:height*0.005,
        borderColor: '#6c7a89',           
        borderWidth: width*0.0005,                 
        borderRadius: 12,
        backgroundColor:"#6c7a89",
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.2,
        shadowRadius: 4,
        elevation: 4,
      },

      tekst:{
        marginBottom:height*0.03,
        fontSize:18,
        fontWeight:'500',
        color:"#fff"
      },

      uredajiGumb: {
        borderRadius: 15,
        paddingHorizontal: width * 0.12,
        paddingVertical: height * 0.015,
        marginTop: height * 0.01,
        marginBottom: height * 0.01,
        backgroundColor: '#6c88a6',
        borderColor: '#b0bec5',
        borderWidth: 1,
        alignItems: 'center',
        alignSelf: 'center',
      },

      uredajiGumbTekst:{
        color:'#fff',
        padding:width*0.003,
        fontSize:15,
        fontWeight:"700"
      },

      lokacijaBlok:{
        paddingHorizontal: width *0.03,
        paddingVertical: height*0.05,
        margin:height*0.005,
        borderColor: '#6c7a89',           
        borderWidth: width*0.0005,                 
        borderRadius: 12,
        backgroundColor:"#6c7a89",
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.2,
        shadowRadius: 4,
        elevation: 4,

      },

      izbornik:{
        flexDirection: 'column',      
        alignItems: 'center',       
        justifyContent: 'space-between',
        padding:width*0.009,
        margin:height*0.005,
        marginVertical:height*0.02,
      },

      karta: {
        width: '100%',
        height: 300,
        borderRadius: 12,
        overflow: Platform.OS === 'ios' ? 'hidden' : 'visible',
      },

      select:{ 
        backgroundColor:"#ecf0f1",
        borderWidth: 1,
        color:"#000",
        paddingHorizontal:width*0.2,
        paddingVertical:height*0.01,
        borderRadius:16,
        fontSize: 16,
        borderColor: '#b0bec5'
      },

      menuConteiner:{
        
        position:'absolute',
        top:50,
        right:15,
        backgroundColor:'#dfe6e9',
        borderRadius: 10,
        borderWidth:1,
        borderColor:"#333",
        paddingHorizontal:width*0.2,
        paddingVertical:height*0.09,
        elevation: 5,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.2,
        shadowRadius: 4,
        zIndex: 10,
        alignItems: 'flex-start', 
        justifyContent: 'center',
      },
    
      menuItems: {
        paddingVertical: height * 0.015,    
        paddingHorizontal: width * 0.05,    
        marginBottom: height * 0.01,        
        backgroundColor: '#aabbee',         
        borderRadius: 8,
      },

      menuText: {
        color: '#000',         
        fontSize: 16,
        fontWeight: '600',
      },
      
      dateButton: {
        backgroundColor: '#ecf0f1',
        paddingVertical: height * 0.015,
        paddingHorizontal: width * 0.05,
        borderRadius: 10,
        borderWidth: 1,
        borderColor: '#b0bec5',
        marginHorizontal: width * 0.02,
        marginBottom: height * 0.015,
        alignItems: 'center',
      },
      
      dateButtonText: {
        color: '#2f80ed',
        fontSize: 16,
        fontWeight: '600',
      },
      
      fullscreenMapContainer: {
        position: 'absolute',
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        zIndex: 100,
        backgroundColor: '#e5e8ea',
      },
      
      zatvoriGumb: {
        position: 'absolute',
        top: 40,
        right: 20,
        zIndex: 101,
        backgroundColor: '#dfe6e9',
        paddingVertical: height * 0.008,
        paddingHorizontal: width * 0.04,
        borderRadius: 10,
        borderWidth: 1,
        borderColor: '#b0bec5',
      },
      
      zatvoriTekst: {
        fontSize: 16,
        color: '#2f80ed',
        fontWeight: '600',
      }
})