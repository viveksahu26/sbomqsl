# discuss about sbom

- SBOMs comes up with 2 format SPDX and CycloneDX.
- Both have different structures, but contains most of the similar fields.
- For SBOM we need those fields.
- For example, SPDX and CycloneDX both have SBOM name, Namespace, version, creation time, Format, Spec types, etc but at different fields.
- In order to get it we need a common interface which provide these all fields irrespective of how they parse it from their structures.
- So, let's define a spec interface:
  - Spec interface defines a set of methods that any type implementing this interface must provide
-  
