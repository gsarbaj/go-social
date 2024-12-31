"use server"

import {redirect} from "next/navigation";
import {headers} from "next/headers";

export const handleConfirm = async () => {

    function getQueryString(url: string): string {
        const queryStart = url.indexOf('?');
        if (queryStart === -1) {
            return ''; // Return an empty string if no query string is found
        }
        return url.substring(queryStart + 1);
    }

    const headersList = await headers()
    const token = getQueryString(headersList.get('referer') || "")

    const API_URL = process.env.API_URL

    const response = await fetch(`${API_URL}/users/activate/${token}`, {
        method: 'PUT'
    })


    if (response.ok) {
        console.log("Failed to activate user")
        redirect("/")
    } else {
        console.log(response.status, response.statusText)
    }
}


// created on 31/12/2024 16:06