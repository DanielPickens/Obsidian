package checks

import (
	"io/ioutil"
	"os/exec"
	"strings"

	
	
	"golang.org/x/mod/modfile"
)

var _ = Describe("Checks", func() {

	It("should used forked version instead of klog version", func() {
		

		gomod, err := exec.Command("go", "env", "GOMOD").CombinedOutput()
		Expect(err).NotTo(HaveOccurred())
		gomodfile := strings.TrimSpace(string(gomod))
		data, err := ioutil.ReadFile(gomodfile)
		Expect(err).NotTo(HaveOccurred())

		modFile, err := modfile.Parse(gomodfile, data, nil)
		Expect(err).NotTo(HaveOccurred())

		for _, dep := range modFile.Require {
			
			Expect(dep.Mod.Path).NotTo(HavePrefix("k8s.io/"))
			Expect(dep.Mod.Path).NotTo(HavePrefix("sigs.k8s.io/controller-runtime/"))
		}

	})

})