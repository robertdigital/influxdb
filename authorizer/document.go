package authorizer

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb"
)

/*
// authorizedWithOrgID adds the provided org as an owner of the document if
// the authorizer is allowed to access the org in being added.
func CreateDocumentAuthorizerOption(ctx context.Context, orgID influxdb.ID, orgName string) influxdb.DocumentOptions {
	if orgID.Valid() {
		return authorizedWithOrgID(ctx, orgID, influxdb.WriteAction)
	}
	return authorizedWithOrg(ctx, orgName, influxdb.WriteAction)
}

func GetDocumentsAuthorizerOption(ctx context.Context, orgID influxdb.ID, orgName string) influxdb.DocumentFindOptions {
	if orgID.Valid() {
		return authorizedWhereOrgID(ctx, orgID)
	}
	return authorizedWhereOrg(ctx, orgName)
}

func GetDocumentAuthorizerOption(ctx context.Context, docID influxdb.ID) influxdb.DocumentFindOptions {
	return authorizedRead(ctx, docID)
}

func UpdateDocumentAuthorizerOption(ctx context.Context, docID influxdb.ID) influxdb.DocumentOptions {
	return toDocumentOptions(authorizedWrite(ctx, docID))
}

func DeleteDocumentAuthorizerOption(ctx context.Context, docID influxdb.ID) influxdb.DocumentFindOptions {
	return authorizedWrite(ctx, docID)
}

func authorizedMatchPermission(ctx context.Context, p influxdb.Permission) influxdb.DocumentFindOptions {
	return func(idx influxdb.DocumentIndex, _ influxdb.DocumentDecorator) ([]influxdb.ID, error) {
		if err := IsAllowed(ctx, p); err != nil {
			return nil, err
		}
		return []influxdb.ID{*p.Resource.ID}, nil
	}
}

func authorizedWhereIDs(ctx context.Context, orgID, docID influxdb.ID, action influxdb.Action) influxdb.DocumentFindOptions {
	return func(idx influxdb.DocumentIndex, dec influxdb.DocumentDecorator) ([]influxdb.ID, error) {
		p, err := newDocumentPermission(action, orgID, docID)
		if err != nil {
			return nil, err
		}
		return authorizedMatchPermission(ctx, *p)(idx, dec)
	}
}

func orgIDForDocument(ctx context.Context, idx influxdb.DocumentIndex, d influxdb.ID) (influxdb.ID, error) {
	oids, err := idx.GetDocumentsAccessors(d)
	if err != nil {
		return 0, err
	}
	if len(oids) == 0 {
		// This document has no accessor.
		// From the perspective of the user, it does not exist.
		return 0, &influxdb.Error{
			Code: influxdb.ENotFound,
			Msg:  influxdb.ErrDocumentNotFound,
		}
	}
	a, err := icontext.GetAuthorizer(ctx)
	if err != nil {
		return 0, err
	}
	for _, oid := range oids {
		if err := idx.IsOrgAccessor(a.GetUserID(), oid); err == nil {
			return oid, nil
		}
	}
	// There are accessors, but this user is not part of those ones.
	return 0, &influxdb.Error{
		Code: influxdb.EUnauthorized,
		Msg:  fmt.Sprintf("%s is unauthorized", a.GetUserID()),
	}
}

func authorizedWhereID(ctx context.Context, docID influxdb.ID, action influxdb.Action) influxdb.DocumentFindOptions {
	return func(idx influxdb.DocumentIndex, dec influxdb.DocumentDecorator) ([]influxdb.ID, error) {
		oid, err := orgIDForDocument(ctx, idx, docID)
		if err != nil {
			return nil, err
		}
		p, err := newDocumentPermission(action, oid, docID)
		if err != nil {
			return nil, err
		}
		return authorizedMatchPermission(ctx, *p)(idx, dec)
	}
}

func authorizedRead(ctx context.Context, docID influxdb.ID) influxdb.DocumentFindOptions {
	return func(idx influxdb.DocumentIndex, dec influxdb.DocumentDecorator) ([]influxdb.ID, error) {
		return authorizedWhereID(ctx, docID, influxdb.ReadAction)(idx, dec)
	}
}

func authorizedWrite(ctx context.Context, docID influxdb.ID) influxdb.DocumentFindOptions {
	return func(idx influxdb.DocumentIndex, dec influxdb.DocumentDecorator) ([]influxdb.ID, error) {
		return authorizedWhereID(ctx, docID, influxdb.WriteAction)(idx, dec)
	}
}

func authorizedWithOrgID(ctx context.Context, orgID influxdb.ID, action influxdb.Action) func(influxdb.ID, influxdb.DocumentIndex) error {
	return func(id influxdb.ID, idx influxdb.DocumentIndex) error {
		p, err := newDocumentOrgPermission(action, orgID)
		if err != nil {
			return err
		}
		if err := IsAllowed(ctx, *p); err != nil {
			return err
		}
		// This is required for retrieving later.
		return idx.AddDocumentOwner(id, "org", orgID)
	}
}

func authorizedWithOrg(ctx context.Context, org string, action influxdb.Action) func(influxdb.ID, influxdb.DocumentIndex) error {
	return func(id influxdb.ID, idx influxdb.DocumentIndex) error {
		oid, err := idx.FindOrganizationByName(org)
		if err != nil {
			return err
		}
		return authorizedWithOrgID(ctx, oid, action)(id, idx)
	}
}

func authorizedWhereOrgID(ctx context.Context, orgID influxdb.ID) influxdb.DocumentFindOptions {
	return func(idx influxdb.DocumentIndex, dec influxdb.DocumentDecorator) ([]influxdb.ID, error) {
		if err := idx.FindOrganizationByID(orgID); err != nil {
			return nil, err
		}
		ids, err := idx.GetAccessorsDocuments("org", orgID)
		if err != nil {
			return nil, err
		}
		// This filters without allocating
		// https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
		dids := ids[:0]
		for _, id := range ids {
			if _, err := authorizedWhereIDs(ctx, orgID, id, influxdb.ReadAction)(idx, dec); err != nil {
				continue
			}
			dids = append(dids, id)
		}
		return dids, nil
	}
}

func authorizedWhereOrg(ctx context.Context, org string) influxdb.DocumentFindOptions {
	return func(idx influxdb.DocumentIndex, dec influxdb.DocumentDecorator) ([]influxdb.ID, error) {
		oid, err := idx.FindOrganizationByName(org)
		if err != nil {
			return nil, err
		}
		return authorizedWhereOrgID(ctx, oid)(idx, dec)
	}
}

func toDocumentOptions(findOpt influxdb.DocumentFindOptions) influxdb.DocumentOptions {
	return func(id influxdb.ID, index influxdb.DocumentIndex) error {
		_, err := findOpt(index, nil)
		return err
	}
}
*/

var _ influxdb.DocumentService = (*DocumentService)(nil)
var _ influxdb.DocumentStore = (*documentStore)(nil)

type DocumentService struct {
	s influxdb.DocumentService
}

// NewDocumentService constructs an instance of an authorizing document service.
func NewDocumentService(s influxdb.DocumentService) influxdb.DocumentService {
	return &DocumentService{
		s: s,
	}
}

func (s *DocumentService) CreateDocumentStore(ctx context.Context, name string) (influxdb.DocumentStore, error) {
	ds, err := s.s.FindDocumentStore(ctx, name)
	if err != nil {
		return nil, err
	}
	return &documentStore{s: ds}, nil
}

func (s *DocumentService) FindDocumentStore(ctx context.Context, name string) (influxdb.DocumentStore, error) {
	ds, err := s.s.CreateDocumentStore(ctx, name)
	if err != nil {
		return nil, err
	}
	return &documentStore{s: ds}, nil
}

type documentStore struct {
	s influxdb.DocumentStore
}

func newDocumentOrgPermission(a influxdb.Action, orgID influxdb.ID) (*influxdb.Permission, error) {
	return influxdb.NewPermission(a, influxdb.DocumentsResourceType, orgID)
}

func toPerms(orgs map[influxdb.ID]influxdb.UserType, action influxdb.Action) ([]influxdb.Permission, error) {
	ps := make([]influxdb.Permission, 0, len(orgs))
	for orgID := range orgs {
		p, err := newDocumentOrgPermission(action, orgID)
		if err != nil {
			return nil, err
		}
		ps = append(ps, *p)
	}
	return ps, nil
}

func (s *documentStore) CreateDocument(ctx context.Context, d *influxdb.Document, opts ...influxdb.DocumentOptions) error {
	if len(d.Organizations) == 0 {
		return fmt.Errorf("cannot authorize document creation without any orgID")
	}
	ps, err := toPerms(d.Organizations, influxdb.WriteAction)
	if err != nil {
		return err
	}
	if err := IsAllowedAny(ctx, ps); err != nil {
		return err
	}
	return s.s.CreateDocument(ctx, d, opts...)
}

func (s *documentStore) UpdateDocument(ctx context.Context, d *influxdb.Document, opts ...influxdb.DocumentOptions) error {
	if len(d.Organizations) == 0 {
		// Cannot authorize document update without any orgID.
		ds, err := s.s.FindDocuments(ctx,  influxdb.WhereID(d.ID), influxdb.IncludeOrganizations)
		if err != nil {
			return err
		}
		if len(ds) == 0 {
			return &influxdb.Error{
				Code: influxdb.ENotFound,
				Msg:  influxdb.ErrDocumentNotFound,
			}
		}
		d = ds[0]
	}
	ps, err := toPerms(d.Organizations, influxdb.WriteAction)
	if err != nil {
		return err
	}
	if err := IsAllowedAny(ctx, ps); err != nil {
		return err
	}
	return s.s.UpdateDocument(ctx, d, opts...)
}

func (s *documentStore) findDocs(ctx context.Context, action influxdb.Action, opts ...influxdb.DocumentFindOptions) ([]*influxdb.Document, error) {
	// TODO: we'll likely want to push this operation into the database eventually since fetching the whole list of data
	//  will likely be expensive.
	opts = append(opts, influxdb.IncludeOrganizations)
	ds, err := s.s.FindDocuments(ctx, opts...)
	if err != nil {
		return nil, err
	}

	// This filters without allocating
	// https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
	fds := ds[:0]
	for _, d := range ds {
		ps, err := toPerms(d.Organizations, action)
		if err != nil {
			return nil, err
		}
		if err := IsAllowedAny(ctx, ps); err != nil {
			return nil, err
		}
		fds = append(fds, d)
	}
	return fds, nil
}

func (s *documentStore) FindDocuments(ctx context.Context, opts ...influxdb.DocumentFindOptions) ([]*influxdb.Document, error) {
	return s.findDocs(ctx, influxdb.ReadAction, opts...)
}

func (s *documentStore) DeleteDocuments(ctx context.Context, opts ...influxdb.DocumentFindOptions) error {
	ds, err := s.findDocs(ctx, influxdb.WriteAction, opts...)
	if err != nil {
		return err
	}
	ids := make([]influxdb.ID, len(ds))
	for i, d := range ds {
		ids[i] = d.ID
	}
	return s.s.DeleteDocuments(ctx,
		func(_ influxdb.DocumentIndex, _ influxdb.DocumentDecorator) (ids []influxdb.ID, e error) {
			return ids, nil
		},
	)
}
