import { Injectable } from "@angular/core";

@Injectable({
	providedIn: "root",
})
export class LoginService {
	submitApplication(login: string, password: string) {
		console.log(
			`Login application received: login : ${login}, password: ${password}.`,
		);
	}
}
