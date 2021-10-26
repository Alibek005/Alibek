// Code generated by go-bindata.
// sources:
// i18n/resources/de_DE.all.json
// i18n/resources/en_US.all.json
// i18n/resources/es_ES.all.json
// i18n/resources/fr_FR.all.json
// i18n/resources/it_IT.all.json
// i18n/resources/ja_JP.all.json
// i18n/resources/ko_KR.all.json
// i18n/resources/pt_BR.all.json
// i18n/resources/zh_Hans.all.json
// i18n/resources/zh_Hant.all.json
// DO NOT EDIT!

package resources

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _i18nResourcesDe_deAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\nGeben Sie eine Zahl ein."
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "Beim Erstellen der Protokolldatei '{{.Path}}' ist ein Fehler aufgetreten:\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "Bei der Anforderung zum Erstellen eines Speicherauszugs ist ein Fehler aufgetreten:\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "Bei der Antwort bezüglich der Erstellung eines Speicherauszugs ist ein Fehler aufgetreten:\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "Verstrichen:"
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "FEHLGESCHLAGEN"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "Ungültiges Token: "
  },
  {
    "id": "OK",
    "translation": "OK"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "Geben Sie 'j', 'n', 'ja' oder 'nein' ein."
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "Geben Sie eine Zahl zwischen 1 und {{.Count}} ein."
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "Geben Sie eine gültige Gleitkommazahl ein."
  },
  {
    "id": "Please enter a valid number.",
    "translation": "Geben Sie eine gültige Zahl ein."
  },
  {
    "id": "Please enter value.",
    "translation": "Geben Sie einen Wert ein."
  },
  {
    "id": "REQUEST:",
    "translation": "ANFORDERUNG:"
  },
  {
    "id": "RESPONSE:",
    "translation": "ANTWORT:"
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "Fehler auf dem fernen Server. Statuscode: {{.StatusCode}}, Fehlercode: {{.ErrorCode}}, Nachricht: {{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "Speichern der Plug-in-Konfiguration nicht möglich: "
  }
]`)

func i18nResourcesDe_deAllJsonBytes() ([]byte, error) {
	return _i18nResourcesDe_deAllJson, nil
}

func i18nResourcesDe_deAllJson() (*asset, error) {
	bytes, err := i18nResourcesDe_deAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/de_DE.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesEn_usAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\nEnter a number"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "An error occurred while dumping request:\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "An error occurred while dumping response:\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "Elapsed:"
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "FAILED"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "Invalid token: "
  },
  {
    "id": "OK",
    "translation": "OK"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "Please enter 'y', 'n', 'yes' or 'no'."
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "Please enter a number between 1 to {{.Count}}."
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "Please enter a valid floating number."
  },
  {
    "id": "Please enter a valid number.",
    "translation": "Please enter a valid number."
  },
  {
    "id": "Please enter value.",
    "translation": "Please enter value."
  },
  {
    "id": "REQUEST:",
    "translation": "REQUEST:"
  },
  {
    "id": "RESPONSE:",
    "translation": "RESPONSE:"
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "Unable to save plugin config: "
  },
  {
    "id": "Session inactive: ",
    "translation": "Session inactive: "
  }
]
`)

func i18nResourcesEn_usAllJsonBytes() ([]byte, error) {
	return _i18nResourcesEn_usAllJson, nil
}

func i18nResourcesEn_usAllJson() (*asset, error) {
	bytes, err := i18nResourcesEn_usAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/en_US.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesEs_esAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\nEscriba un número"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "Se ha producido un error al crear el archivo de registro '{{.Path}}':\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "Se ha producido un error al volcar la solicitud:\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "Se ha producido un error al volcar la respuesta:\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "Transcurrido:"
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "ERROR"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "Señal no válida: "
  },
  {
    "id": "OK",
    "translation": "Correcto"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "Especifique 'y', 'n', 'yes' o 'no'."
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "Especifique un número entre 1 y {{.Count}}."
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "Especifique un número flotante válido."
  },
  {
    "id": "Please enter a valid number.",
    "translation": "Especifique un número válido."
  },
  {
    "id": "Please enter value.",
    "translation": "Especifique un valor."
  },
  {
    "id": "REQUEST:",
    "translation": "SOLICITUD:"
  },
  {
    "id": "RESPONSE:",
    "translation": "RESPUESTA:"
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "Error del servidor remoto. Código de estado: {{.StatusCode}}, código de error: {{.ErrorCode}}, mensaje: {{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "No se ha podido guardar la configuración del plugin:"
  }
]`)

func i18nResourcesEs_esAllJsonBytes() ([]byte, error) {
	return _i18nResourcesEs_esAllJson, nil
}

func i18nResourcesEs_esAllJson() (*asset, error) {
	bytes, err := i18nResourcesEs_esAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/es_ES.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesFr_frAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\nEntrez un nombre"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "Erreur lors de la création du fichier journal '{{.Path}}':\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "Erreur lors de la demande de vidage :\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "Erreur lors de la réponse de vidage :\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "Ecoulé :"
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "ECHEC"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "Jeton non valide : "
  },
  {
    "id": "OK",
    "translation": "OK"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "L'entrée doit être 'y', 'n', 'yes' ou 'no'."
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "Entrez un nombre entre 1 et {{.Count}}."
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "Entrez un nombre flottant valide."
  },
  {
    "id": "Please enter a valid number.",
    "translation": "Entrez un nombre valide."
  },
  {
    "id": "Please enter value.",
    "translation": "Entrez une valeur."
  },
  {
    "id": "REQUEST:",
    "translation": "DEMANDE :"
  },
  {
    "id": "RESPONSE:",
    "translation": "REPONSE :"
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "Erreur du serveur distant. Code de statut : {{.StatusCode}}, code d'erreur : {{.ErrorCode}}, message : {{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "Impossible d'enregistrer la configuration du plug-in : "
  }
]`)

func i18nResourcesFr_frAllJsonBytes() ([]byte, error) {
	return _i18nResourcesFr_frAllJson, nil
}

func i18nResourcesFr_frAllJson() (*asset, error) {
	bytes, err := i18nResourcesFr_frAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/fr_FR.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesIt_itAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\nImmetti un numero"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "Si è verificato un errore durante la creazione del file di log '{{.Path}}':\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "Si è verificato un errore durante il dump della richiesta:\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "Si è verificato un errore durante il dump della risposta:\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "Trascorso:"
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "NON RIUSCITO"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "Token non valido: "
  },
  {
    "id": "OK",
    "translation": "OK"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "Immetti 's', 'n', 'sì' o 'no'."
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "Immetti un numero compreso tra 1 e {{.Count}}."
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "Immetti un numero decimale valido."
  },
  {
    "id": "Please enter a valid number.",
    "translation": "Immetti un numero valido."
  },
  {
    "id": "Please enter value.",
    "translation": "Immetti un valore."
  },
  {
    "id": "REQUEST:",
    "translation": "RICHIESTA:"
  },
  {
    "id": "RESPONSE:",
    "translation": "RISPOSTA:"
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "Errore server remoto. Codice di stato: {{.StatusCode}}, codice di errore: {{.ErrorCode}}, messaggio: {{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "Impossibile salvare la configurazione del plug-in: "
  }
]`)

func i18nResourcesIt_itAllJsonBytes() ([]byte, error) {
	return _i18nResourcesIt_itAllJson, nil
}

func i18nResourcesIt_itAllJson() (*asset, error) {
	bytes, err := i18nResourcesIt_itAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/it_IT.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesJa_jpAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\n数値を入力してください"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "ログ・ファイル '{{.Path}}' を作成中にエラーが発生しました:\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "要求のダンプ中にエラーが発生しました:\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "応答のダンプ中にエラーが発生しました:\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "経過:"
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "失敗"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "トークンが無効です: "
  },
  {
    "id": "OK",
    "translation": "OK"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "「y」、「n」、「yes」、または「no」を入力してください。"
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "1 から {{.Count}} までの数値を入力してください。"
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "有効な浮動小数点数を入力してください。"
  },
  {
    "id": "Please enter a valid number.",
    "translation": "有効な数値を入力してください。"
  },
  {
    "id": "Please enter value.",
    "translation": "値を入力してください。"
  },
  {
    "id": "REQUEST:",
    "translation": "要求:"
  },
  {
    "id": "RESPONSE:",
    "translation": "応答:"
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "リモート・サーバー・エラー。 状況コード: {{.StatusCode}}、エラー・コード: {{.ErrorCode}}、メッセージ: {{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "プラグイン構成を保存できません: "
  }
]`)

func i18nResourcesJa_jpAllJsonBytes() ([]byte, error) {
	return _i18nResourcesJa_jpAllJson, nil
}

func i18nResourcesJa_jpAllJson() (*asset, error) {
	bytes, err := i18nResourcesJa_jpAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/ja_JP.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesKo_krAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\n번호 입력"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "'{{.Path}}' 로그 파일을 작성할 때 다음 오류가 발생했습니다. \n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "요청을 덤프할 때 다음 오류가 발생했습니다. \n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "응답을 덤프할 때 다음 오류가 발생했습니다. \n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "경과 시간:"
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "실패"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "올바르지 않은 토큰: "
  },
  {
    "id": "OK",
    "translation": "확인"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "'y', 'n', '예' 또는 '아니오'를 입력하십시오."
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "1 - {{.Count}} 사이의 수를 입력하십시오."
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "올바른 float 수를 입력하십시오."
  },
  {
    "id": "Please enter a valid number.",
    "translation": "올바른 수를 입력하십시오."
  },
  {
    "id": "Please enter value.",
    "translation": "값을 입력하십시오."
  },
  {
    "id": "REQUEST:",
    "translation": "요청:"
  },
  {
    "id": "RESPONSE:",
    "translation": "응답:"
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "원격 서버 오류가 발생했습니다. 상태 코드: {{.StatusCode}}, 오류 코드: {{.ErrorCode}}, 메시지: {{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "플러그인 구성을 저장할 수 없음:"
  }
]`)

func i18nResourcesKo_krAllJsonBytes() ([]byte, error) {
	return _i18nResourcesKo_krAllJson, nil
}

func i18nResourcesKo_krAllJson() (*asset, error) {
	bytes, err := i18nResourcesKo_krAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/ko_KR.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesPt_brAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\nInsira um número"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "Ocorreu um erro ao criar o arquivo de log '{{.Path}}':\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "Ocorreu um erro ao fazer dump da solicitação:\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "Ocorreu um erro ao fazer dump da resposta:\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "Decorrido:"
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "COM FALHA"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "Token inválido: "
  },
  {
    "id": "OK",
    "translation": "OK"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "Insira 'y', 'n', 'yes' ou 'no'."
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "Insira um número entre 1 e {{.Count}}."
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "Insira um número flutuante válido."
  },
  {
    "id": "Please enter a valid number.",
    "translation": "Digite um número válido."
  },
  {
    "id": "Please enter value.",
    "translation": "Insira um valor."
  },
  {
    "id": "REQUEST:",
    "translation": "SOLICITAÇÃO:"
  },
  {
    "id": "RESPONSE:",
    "translation": "RESPOSTA:"
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "Erro do servidor remoto. Código de status: {{.StatusCode}}, código de erro: {{.ErrorCode}}, mensagem: {{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "Não é possível salvar a configuração do plug-in: "
  }
]`)

func i18nResourcesPt_brAllJsonBytes() ([]byte, error) {
	return _i18nResourcesPt_brAllJson, nil
}

func i18nResourcesPt_brAllJson() (*asset, error) {
	bytes, err := i18nResourcesPt_brAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/pt_BR.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesZh_hansAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\n请输入数字"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "创建日志文件“{{.Path}}”时发生错误：\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "转储请求时发生错误：\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "转储响应时发生错误：\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "经过时长："
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "失败"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "令牌无效："
  },
  {
    "id": "OK",
    "translation": "确定"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "请输入“y”、“n”、“yes”或“no”。"
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "请输入 1 到 {{.Count}} 之间的数字。"
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "请输入有效的浮点数。"
  },
  {
    "id": "Please enter a valid number.",
    "translation": "请输入有效的数字。"
  },
  {
    "id": "Please enter value.",
    "translation": "请输入有效的值。"
  },
  {
    "id": "REQUEST:",
    "translation": "请求: "
  },
  {
    "id": "RESPONSE:",
    "translation": "响应: "
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "远程服务器错误。状态码：{{.StatusCode}}，错误代码：{{.ErrorCode}}，消息：{{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "无法保存插件配置："
  }
]`)

func i18nResourcesZh_hansAllJsonBytes() ([]byte, error) {
	return _i18nResourcesZh_hansAllJson, nil
}

func i18nResourcesZh_hansAllJson() (*asset, error) {
	bytes, err := i18nResourcesZh_hansAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/zh_Hans.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesZh_hantAllJson = []byte(`[
  {
    "id": "\nEnter a number",
    "translation": "\n請輸入數字"
  },
  {
    "id": "An error occurred when creating log file '{{.Path}}':\n{{.Error}}\n\n",
    "translation": "建立日誌檔 '{{.Path}}' 時發生錯誤：\n{{.Error}}\n\n"
  },
  {
    "id": "An error occurred while dumping request:\n{{.Error}}\n",
    "translation": "傾出要求時發生錯誤：\n{{.Error}}\n"
  },
  {
    "id": "An error occurred while dumping response:\n{{.Error}}\n",
    "translation": "傾出回應時發生錯誤：\n{{.Error}}\n"
  },
  {
    "id": "Could not read from input: ",
    "translation": "Could not read from input: "
  },
  {
    "id": "Elapsed:",
    "translation": "經歷時間："
  },
  {
    "id": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "External authentication failed. Error code: {{.ErrorCode}}, message: {{.Message}}"
  },
  {
    "id": "FAILED",
    "translation": "失敗"
  },
  {
    "id": "Invalid grant type: ",
    "translation": "Invalid grant type: "
  },
  {
    "id": "Invalid token: ",
    "translation": "無效的記號："
  },
  {
    "id": "OK",
    "translation": "確定"
  },
  {
    "id": "Please enter 'y', 'n', 'yes' or 'no'.",
    "translation": "請輸入 'y'、'n'、'yes' 或 'no'。"
  },
  {
    "id": "Please enter a number between 1 to {{.Count}}.",
    "translation": "請輸入 1 到 {{.Count}} 之間的數字。"
  },
  {
    "id": "Please enter a valid floating number.",
    "translation": "請輸入有效的浮點數。"
  },
  {
    "id": "Please enter a valid number.",
    "translation": "請輸入有效的數字。"
  },
  {
    "id": "Please enter value.",
    "translation": "請輸入值。"
  },
  {
    "id": "REQUEST:",
    "translation": "要求："
  },
  {
    "id": "RESPONSE:",
    "translation": "回應："
  },
  {
    "id": "Remote server error. Status code: {{.StatusCode}}, error code: {{.ErrorCode}}, message: {{.Message}}",
    "translation": "遠端伺服器錯誤。狀態碼：{{.StatusCode}}，錯誤碼：{{.ErrorCode}}，訊息：{{.Message}}"
  },
  {
    "id": "Unable to save plugin config: ",
    "translation": "無法儲存外掛程式配置："
  }
]`)

func i18nResourcesZh_hantAllJsonBytes() ([]byte, error) {
	return _i18nResourcesZh_hantAllJson, nil
}

func i18nResourcesZh_hantAllJson() (*asset, error) {
	bytes, err := i18nResourcesZh_hantAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/zh_Hant.all.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"i18n/resources/de_DE.all.json":   i18nResourcesDe_deAllJson,
	"i18n/resources/en_US.all.json":   i18nResourcesEn_usAllJson,
	"i18n/resources/es_ES.all.json":   i18nResourcesEs_esAllJson,
	"i18n/resources/fr_FR.all.json":   i18nResourcesFr_frAllJson,
	"i18n/resources/it_IT.all.json":   i18nResourcesIt_itAllJson,
	"i18n/resources/ja_JP.all.json":   i18nResourcesJa_jpAllJson,
	"i18n/resources/ko_KR.all.json":   i18nResourcesKo_krAllJson,
	"i18n/resources/pt_BR.all.json":   i18nResourcesPt_brAllJson,
	"i18n/resources/zh_Hans.all.json": i18nResourcesZh_hansAllJson,
	"i18n/resources/zh_Hant.all.json": i18nResourcesZh_hantAllJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"i18n": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"de_DE.all.json":   &bintree{i18nResourcesDe_deAllJson, map[string]*bintree{}},
			"en_US.all.json":   &bintree{i18nResourcesEn_usAllJson, map[string]*bintree{}},
			"es_ES.all.json":   &bintree{i18nResourcesEs_esAllJson, map[string]*bintree{}},
			"fr_FR.all.json":   &bintree{i18nResourcesFr_frAllJson, map[string]*bintree{}},
			"it_IT.all.json":   &bintree{i18nResourcesIt_itAllJson, map[string]*bintree{}},
			"ja_JP.all.json":   &bintree{i18nResourcesJa_jpAllJson, map[string]*bintree{}},
			"ko_KR.all.json":   &bintree{i18nResourcesKo_krAllJson, map[string]*bintree{}},
			"pt_BR.all.json":   &bintree{i18nResourcesPt_brAllJson, map[string]*bintree{}},
			"zh_Hans.all.json": &bintree{i18nResourcesZh_hansAllJson, map[string]*bintree{}},
			"zh_Hant.all.json": &bintree{i18nResourcesZh_hantAllJson, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
