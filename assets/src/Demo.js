import React, { useState, useEffect } from 'react';
import axios from 'axios'

const Demo = () => {
    const [val, setVal] = useState('loading...');

    useEffect(() => {
        axios.get('/ping').then(resp => {
            console.log(resp);
            setVal(resp.data);
        })
    });

    return <h1>{val}</h1>
};

export default Demo;