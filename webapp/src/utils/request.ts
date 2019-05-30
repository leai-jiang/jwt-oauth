import axios, { AxiosPromise, AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";

interface RequestOption {
	action: string;
	method?: "GET" | "POST";
	headers?: any;
	payload?: any;
}
export const request = (option: RequestOption): AxiosPromise<AxiosResponse> => {
	const request: AxiosRequestConfig = {
		url: option.action,
		method: option.method || "POST",
		headers: option.headers || {},
		data: option.payload || {}
	};

	return axios(request)
		.then((res: AxiosResponse) => {
			return res.data;
		})
		.catch((e: AxiosError) => {
			console.log(e);
		});
};

export default request;
