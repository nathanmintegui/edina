import type { Routes } from "@angular/router";
import { LoginComponent } from "./login/login.component";

export const routes: Routes = [
	{
		path: "login",
		title: "Login Screen",
		component: LoginComponent,
	},
];
