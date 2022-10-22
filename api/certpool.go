package api

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)



type CertPool struct {
	byname map[string]*x509.CertPool

	bydir map[string]*x509.CertPool

	byfile map[string]*x509.CertPool

	bydata map[string]*x509.CertPool

	byurl map[string]*x509.CertPool

	lazy map[string]*x509.CertPool

	lazyCertificate map[string]*tls.Certificate

	haveSum map[string]bool

	SystemCertPool bool


}

type lazyCertificate struct {

	rawSubject []byte

	getCert func() (*tls.Certificate, error)
}

func NewCertPool() *CertPool {
	return &CertPool{
		byname:          make(map[string]*x509.CertPool),
		bydir:           make(map[string]*x509.CertPool),
		byfile:          make(map[string]*x509.CertPool),
		bydata:          make(map[string]*x509.CertPool),
		byurl:           make(map[string]*x509.CertPool),
		lazy:            make(map[string]*x509.CertPool),
		lazyCertificate: make(map[string]*tls.Certificate),
		haveSum:         make(map[string]bool),
	}
}




func (cp *CertPool) AddCert(cert *x509.Certificate,) {
	systemCertPool := cp.SystemCertPool()
	if systemCertPool != nil {
	if cp.SystemCertPool {
		if err := cp.systemCertPool().AddCert(cert); err != nil {
			panic(err)
		}
		return
	}
	cp.bydata[string(cert.Raw)] = x509.NewCertPool()
	cp.bydata[string(cert.Raw)].AddCert(cert)
}

func (cp *CertPool) Subjects() [][]byte {
	res = make([][]byte, 0, len(cp.bydata))
	for _, pool := range cp.bydata {
		for _, cert := range pool.Subjects() {
			res = append(res, cert)
		}
	}
	return res
}

func (cp *CertPool) Certificates() []*x509.Certificate {
	res = make([]*x509.Certificate, 0, len(cp.bydata))
	for _, pool := range cp.bydata {
		for _, cert := range pool.Subjects() {
			res = append(res, cert)
		}
	}
	return res
}

func (cp *CertPool) GetCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	if cert, ok := cp.lazyCertificate[clientHello.ServerName]; ok {
		return cert, nil
	}
}
}




func (cp *CertPool) SystemCertPool() *x509.CertPool {
	if cp.SystemCertPool {
		return cp.systemCertPool()
	}
	return nil
}

func (cp *CertPool) systemCertPool() *x509.CertPool {
	if cp.SystemCertPool {
		return x509.SystemCertPool()
	}
	return nil
}

func (cp *CertPool) GetCertPool(name string) (*x509.CertPool, error) {
	if pool, ok := cp.byname[name]; ok {
		return pool, nil
	}
	if pool, ok := cp.bydir[name]; ok {
		return pool, nil
	}
	if pool, ok := cp.byfile[name]; ok {
		return pool, nil
	}
	if pool, ok := cp.bydata[name]; ok {
		return pool, nil
	}
	if pool, ok := cp.byurl[name]; ok {
		return pool, nil
	}
	if pool, ok := cp.lazy[name]; ok {
		return pool, nil
	}
	return nil, fmt.Errorf("cert pool %q not found", name)
}




func (cp *CertPool) mutex() *x509.CertPool {
	var mu = new(sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	return cp

}

func worker() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Concurrency waitgroup for goroutine")
		}()
	}
	wg.Wait()
	fmt.Println("All goroutines complete")
}


	


//It is not safe to call this concurrently with any other method on this type so we must add a mutex to protect it.

func Mutex() {
	var mu = new(sync.Mutex)
	var sync = new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		sync.Add(1)
		go func() {
			defer sync.Done()

	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Mutex for goroutine")
		}()
}
sync.Wait()
fmt.Println("All goroutines complete")
}




//GetClient returns a http client with the given cert pool and the given cert
func GetClient() (*http.Client, error) {
	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	// Read in the cert file
	certs, err := filepath.Glob("/etc/ssl/certs/*.pem")
	if err != nil {
		return nil, err
	}
	for _, cert := range certs {
		cert, err := ioutil.ReadFile(cert)
		if err != nil {
			return nil, err
		}
		// Append our cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(cert); !ok {
			fmt.Println("No certs appended, using system certs only")
		}
	}
	// Trust the augmented cert pool in our client
	config := &tls.Config{
		RootCAs: rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}
	return client, nil
}



func CertPool() (*x509.CertPool, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	if certPool == nil {
		certPool = x509.NewCertPool()
	}

	certFiles, err := filepath.Glob("/etc/ssl/certs/*.pem")
	if err != nil {
		return nil, err
	}

	for _, certFile := range certFiles {
		cert, err := ioutil.ReadFile(certFile)
		if err != nil {
			return nil, err
		}
		certPool.AppendCertsFromPEM(cert)
	}

	return certPool, nil
}

func GetClientWithCertPool() (*http.Client, error) {

	certPool, err := CertPool()
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &http.Client{
		Transport: transport,
	}, nil
}


func (cp *CertPool) FoundData(cert *x509.Certificate) bool {
	if _, ok := cp.bydata[string(cert.Raw)]; ok {
		return true
	
	var mu = new(sync.Mutex)

	for _, pool := range cp.bydata {
		mu.Lock()
		defer mu.Unlock()
		for _, cert := range pool.Subjects() {
			if cert.Equal(cert) {
				return true
			}
		}
	}
	return false
}
}

func (cp *CertPool) AddData(cert *x509.Certificate) {
	if cp.FoundData(cert) {
	if _, ok := cp.bydata[data]; ok {
		return 
	}else {
		pool := x509.NewCertPool()
		pool.AddCert(cert)
		cp.bydata[data] = pool
	}
}

func (cp *CertPool) AddFile(certFile string) error {
	cert, err := ioutil.ReadFile(certFile)
	if err != nil {
		return err
	}
}
}

