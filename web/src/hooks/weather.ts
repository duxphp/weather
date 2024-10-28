import {useState} from "react";

const URI = 'http://0.0.0.0:8800/api/weather'


export const useWeather = () => {

    const [weather, setWeather] = useState<Record<string, any>>()
    const [loading, setLoading] = useState(false)

    const get = (city: string) => {
        setLoading(true)
        fetch(`${URI}/${city}`).then(res => res.json()).then(data => {
            if (data.status !== 200) {
                console.error(data.message)
                setWeather(undefined)
            }
            setWeather(data.data[0])
        }).catch(err => {
            console.log(err)
            alert(err.error)
            setWeather(undefined)
        }).finally(() => {
            setTimeout(() => setLoading(false), 2000)
        })
    }

    return {
        weather,
        loading,
        get
    }
}