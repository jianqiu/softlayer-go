package services_test

import (
	"errors"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclient "github.com/maximilien/softlayer-go/client"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

const (
	TEST_NOTES_PREFIX = "TEST:softlayer-go"
)

var _ = Describe("SoftLayer Services", func() {
	var (
		username, apiKey string
		err              error

		client softlayer.Client

		accountService      softlayer.SoftLayer_Account_Service
		virtualGuestService softlayer.SoftLayer_Virtual_Guest_Service
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		client = slclient.NewSoftLayerClient(username, apiKey)
		Expect(client).ToNot(BeNil())

		accountService, err = client.GetSoftLayer_Account_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(accountService).ToNot(BeNil())

		virtualGuestService, err = client.GetSoftLayer_Virtual_Guest_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(virtualGuestService).ToNot(BeNil())
	})

	Context("uses SoftLayer_Account to list current virtual: disk images, guests, ssh keys, and network storage", func() {
		It("returns an array of SoftLayer_Virtual_Guest disk images", func() {
			virtualDiskImages, err := accountService.GetVirtualDiskImages()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(virtualDiskImages)).To(BeNumerically(">=", 0))
		})

		It("returns an array of SoftLayer_Virtual_Guest objects", func() {
			virtualGuests, err := accountService.GetVirtualGuests()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(virtualGuests)).To(BeNumerically(">=", 0))
		})

		It("returns an array of SoftLayer_Virtual_Guest network storage", func() {
			networkStorageArray, err := accountService.GetNetworkStorage()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(networkStorageArray)).To(BeNumerically(">=", 0))
		})

		It("returns an array of SoftLayer_Ssh_Keys objects", func() {
			sshKeys, err := accountService.GetSshKeys()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(sshKeys)).To(BeNumerically(">=", 0))
		})
	})

	XContext("uses SoftLayer_Account to create and then delete a an ssh key", func() {
		It("creates the ssh key and verify it is present", func() {
			Expect(false).To(BeTrue())
		})

		It("deletes the ssh key", func() {
			Expect(false).To(BeTrue())
		})
	})

	XContext("uses SoftLayer_Account to create and then delete a virtual guest instance", func() {
		It("creates the virtual guest instance and waits for it to be active", func() {
			Expect(false).To(BeTrue())
		})

		It("deletes the virtual guest instance if it is running", func() {
			Expect(false).To(BeTrue())
		})
	})

	XContext("uses SoftLayer_Account to create a new instance and network storage and attach them", func() {
		It("creates the virtual guest instance and waits for it to be active", func() {
			Expect(false).To(BeTrue())
		})

		It("creates the disk storage and attaches it to the instance", func() {
			Expect(false).To(BeTrue())
		})

		It("deletes the virtual guest instance if it is running", func() {
			Expect(false).To(BeTrue())
		})

		It("detaches and deletes the network storage if available", func() {
			Expect(false).To(BeTrue())
		})
	})
})

func findTestVirtualGuests() ([]datatypes.SoftLayer_Virtual_Guest, error) {
	accountService, err := createAccountService()
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Guest{}, err
	}

	virtualGuests, err := accountService.GetVirtualGuests()
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Guest{}, err
	}

	testVirtualGuests := []datatypes.SoftLayer_Virtual_Guest{}
	for _, vGuest := range virtualGuests {
		if strings.Contains(vGuest.Notes, TEST_NOTES_PREFIX) {
			testVirtualGuests = append(testVirtualGuests, vGuest)
		}
	}

	return testVirtualGuests, nil
}

func findTestVirtualDiskImages() ([]datatypes.SoftLayer_Virtual_Disk_Image, error) {
	accountService, err := createAccountService()
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Disk_Image{}, err
	}

	virtualDiskImages, err := accountService.GetVirtualDiskImages()
	if err != nil {
		return []datatypes.SoftLayer_Virtual_Disk_Image{}, err
	}

	testVirtualDiskImages := []datatypes.SoftLayer_Virtual_Disk_Image{}
	for _, vDI := range virtualDiskImages {
		if strings.Contains(vDI.Description, TEST_NOTES_PREFIX) {
			testVirtualDiskImages = append(testVirtualDiskImages, vDI)
		}
	}

	return testVirtualDiskImages, nil
}

func findTestNetworkStorage() ([]datatypes.SoftLayer_Network_Storage, error) {
	accountService, err := createAccountService()
	if err != nil {
		return []datatypes.SoftLayer_Network_Storage{}, err
	}

	networkStorageArray, err := accountService.GetNetworkStorage()
	if err != nil {
		return []datatypes.SoftLayer_Network_Storage{}, err
	}

	testNetworkStorageArray := []datatypes.SoftLayer_Network_Storage{}
	for _, storage := range networkStorageArray {
		if strings.Contains(storage.Notes, TEST_NOTES_PREFIX) {
			testNetworkStorageArray = append(testNetworkStorageArray, storage)
		}
	}

	return testNetworkStorageArray, nil
}

func findTestSshKeys() ([]datatypes.SoftLayer_Ssh_Key, error) {
	accountService, err := createAccountService()
	if err != nil {
		return []datatypes.SoftLayer_Ssh_Key{}, err
	}

	sshKeys, err := accountService.GetSshKeys()
	if err != nil {
		return []datatypes.SoftLayer_Ssh_Key{}, err
	}

	testSshKeys := []datatypes.SoftLayer_Ssh_Key{}
	for _, key := range sshKeys {
		if strings.Contains(key.Notes, TEST_NOTES_PREFIX) {
			testSshKeys = append(testSshKeys, key)
		}
	}

	return testSshKeys, nil
}

func getUsernameAndApiKey() (string, string, error) {
	username := os.Getenv("SL_USERNAME")
	if username == "" {
		return "", "", errors.New("SL_USERNAME environment must be set")
	}

	apiKey := os.Getenv("SL_API_KEY")
	if apiKey == "" {
		return username, "", errors.New("SL_API_KEY environment must be set")
	}

	return username, apiKey, nil
}

func createAccountService() (softlayer.SoftLayer_Account_Service, error) {
	username, apiKey, err := getUsernameAndApiKey()
	if err != nil {
		return nil, err
	}

	client := slclient.NewSoftLayerClient(username, apiKey)
	accountService, err := client.GetSoftLayer_Account_Service()
	if err != nil {
		return nil, err
	}

	return accountService, nil
}

func createVirtualGuestService() (softlayer.SoftLayer_Virtual_Guest_Service, error) {
	username, apiKey, err := getUsernameAndApiKey()
	if err != nil {
		return nil, err
	}

	client := slclient.NewSoftLayerClient(username, apiKey)
	virtualGuestService, err := client.GetSoftLayer_Virtual_Guest_Service()
	if err != nil {
		return nil, err
	}

	return virtualGuestService, nil
}

func createSshKeyService() (softlayer.SoftLayer_Ssh_Key_Service, error) {
	username, apiKey, err := getUsernameAndApiKey()
	if err != nil {
		return nil, err
	}

	client := slclient.NewSoftLayerClient(username, apiKey)
	sshKeyService, err := client.GetSoftLayer_Ssh_Key_Service()
	if err != nil {
		return nil, err
	}

	return sshKeyService, nil
}
