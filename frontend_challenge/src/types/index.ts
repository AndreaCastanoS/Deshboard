// Interfaz para los datos traidos de mi db
export interface DataRecord {
  ID: number;
  Age: number;
  Education: string;
  MaritalStatus: string;
  Occupation: string;
  Income: string;
}


export interface JWTPayload {
    MapClaims: {
      eat: number
      iat: number
    }
    session: string
  }