// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPrivatecaCertificateAuthorityIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/privateca.certificateManager",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateAuthorityIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccPrivatecaCertificateAuthorityIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccPrivatecaCertificateAuthorityIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/privateca.certificateManager",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccPrivatecaCertificateAuthorityIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccPrivatecaCertificateAuthorityIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/privateca.certificateManager",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateAuthorityIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccPrivatecaCertificateAuthorityIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccPrivatecaCertificateAuthorityIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  disable_on_delete = true
}

resource "google_privateca_certificate_authority_iam_member" "foo" {
  provider = google-beta
  certificate_authority = google_privateca_certificate_authority.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccPrivatecaCertificateAuthorityIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  disable_on_delete = true
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_privateca_certificate_authority_iam_policy" "foo" {
  provider = google-beta
  certificate_authority = google_privateca_certificate_authority.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPrivatecaCertificateAuthorityIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  disable_on_delete = true
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_privateca_certificate_authority_iam_policy" "foo" {
  provider = google-beta
  certificate_authority = google_privateca_certificate_authority.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPrivatecaCertificateAuthorityIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  disable_on_delete = true
}

resource "google_privateca_certificate_authority_iam_binding" "foo" {
  provider = google-beta
  certificate_authority = google_privateca_certificate_authority.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccPrivatecaCertificateAuthorityIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  disable_on_delete = true
}

resource "google_privateca_certificate_authority_iam_binding" "foo" {
  provider = google-beta
  certificate_authority = google_privateca_certificate_authority.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}